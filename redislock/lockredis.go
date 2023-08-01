package redislock

import (
	"errors"
	"fmt"
	"log"
	"time"
)

// Lock 加锁
func (lock *RedisLock) Lock() error {
	lock.mutex.Lock()
	defer lock.mutex.Unlock()

	result, err := lock.Client.Eval(lock.Context, luaLock, []string{lock.key}, lock.token, lock.lockTimeout.Seconds()).Result()

	if err != nil {
		return fmt.Errorf("failed to acquire lock: %w", err)
	}

	if result != "OK" {
		return errors.New("lock acquisition failed")
	}

	lock.lockCounter++
	if lock.isAutoRenew {
		go lock.autoRenew()
	}
	return nil

}

// UnLock 解锁
func (lock *RedisLock) UnLock() error {
	lock.mutex.Lock()
	defer lock.mutex.Unlock()

	// 可重入锁计数器-1
	if lock.lockCounter > 1 {
		lock.Client.Decr(lock.Context, lock.key)
		lock.lockCounter--
		return nil
	}

	result, err := lock.Client.Eval(lock.Context, luaUnLock, []string{lock.key}, lock.token).Result()

	if err != nil {
		return fmt.Errorf("ailed to release lock: %w", err)
	}

	if result != "OK" {
		return errors.New("lock release failed")
	}

	lock.lockCounter = 0
	return nil
}

// SpinLock 自旋锁
func (lock *RedisLock) SpinLock(timeout time.Duration) error {
	exp := time.Now().Add(timeout)
	for {
		if time.Now().After(exp) {
			return errors.New("spin lock timeout")
		}

		// 加锁成功直接返回
		err := lock.Lock()
		if err == nil {
			return nil
		}

		// 如果加锁失败，则休眠一段时间再尝试
		select {
		case <-lock.Context.Done():
			return lock.Context.Err() // 处理取消操作
		case <-time.After(100 * time.Millisecond):
			// 继续尝试下一轮加锁
		}
	}
}

// Renew 锁手动续期
func (lock *RedisLock) Renew() error {
	lock.mutex.Lock()
	defer lock.mutex.Unlock()

	res, err := lock.Client.Eval(lock.Context, luaRenew, []string{lock.key}, lock.token, lock.lockTimeout.Seconds()).Result()

	if err != nil {
		return fmt.Errorf("failed to renew lock: %s", err)
	}

	if res != "OK" {
		return errors.New("lock renewal failed")
	}

	return nil
}

// 锁自动续期
func (lock *RedisLock) autoRenew() {
	ticker := time.NewTicker(lock.lockTimeout / 2)
	defer ticker.Stop()

	for {
		select {
		case <-lock.Context.Done():
			return
		case <-ticker.C:
			err := lock.Renew()
			if err != nil {
				log.Println("autoRenew failed:", err)
				return
			}
		}
	}
}
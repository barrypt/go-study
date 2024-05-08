package main

import (
	"context"
	"errors"
	"fmt"
	"strings"
	"sync"

	"golang.org/x/sync/errgroup"
)

func  main(){

 cax := context.Background()
    search(cax,"GO")
	coSearch(cax,[]string{"go","c#","python"})
	coGroupSearch(cax,[]string{"go","c#","python"})
}

func search(ctx context.Context, word string) (string, error) {
 select {
 case <-ctx.Done():
  return "", ctx.Err()
 default:
  if strings.ToLower(word) == "go" ||  strings.ToLower(word) == "c#" {
   return "", errors.New("go or C#")
  }
  return fmt.Sprintf("result: %s", word), nil // 模拟结果
 }
}

func coSearch(ctx context.Context, words []string) ([]string, error) {
 ctx, cancel := context.WithCancelCause(ctx)
 defer cancel(nil)

 var (
  wg   = sync.WaitGroup{}
  once = sync.Once{}

  results = make([]string, len(words))
  tokens  = make(chan struct{}, 2)

  err error
 )

 for i, word := range words {
  tokens <- struct{}{}
  wg.Add(1)

  go func(word string, i int) {
   defer func() {
    wg.Done()
    <-tokens
   }()

   result, e := search(ctx, word)
   if e != nil {
    once.Do(func() {
     err = e
     cancel(e)
    })

    return
   }

   results[i] = result
  }(word, i)
 }

 wg.Wait()

 return results, err
}
func coGroupSearch(ctx context.Context, words []string) ([]string, error) {
 g, ctx := errgroup.WithContext(ctx)
 g.SetLimit(10)
 
 results := make([]string, len(words))

 for i, word := range words {
  i, word := i, word

  g.Go(func() error {
   result, err := search(ctx, word)
   if err != nil {
    return err
   }

   results[i] = result
   return nil
  })
 }

 err := g.Wait()

 return results, err

}
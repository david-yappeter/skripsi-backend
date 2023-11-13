package util

import "golang.org/x/sync/errgroup"

func Await(fn func(group *errgroup.Group)) error {
	g := new(errgroup.Group)

	fn(g)

	if err := g.Wait(); err != nil {
		return err
	}

	return nil
}

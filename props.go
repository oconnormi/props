package main

import (
  "os"
  "fmt"

  "github.com/urfave/cli"
  "github.com/oconnormi/properties"

  "github.com/oconnormi/props/version"
)

func main() {
  app := cli.NewApp()
  app.Name = "props"
  app.Version = version.Version
  app.Usage = "modifies property files"
  app.Commands = []cli.Command{
    {
      Name:     "get",
      Aliases:  []string{"g"},
      Usage:    "get the value of a property",
      Action:   func (c *cli.Context) error {
        key := c.Args().Get(0)
        path := c.Args().Get(1)
        p := properties.MustLoadFile(path, properties.UTF8)

        value := p.MustGet(key)
        if value == "" {
          return cli.NewExitError("no such property", 1)
        }
        fmt.Println(value)
        return nil
      },
    },
    {
      Name:     "set",
      Aliases:  []string{"s"},
      Usage:    "set the value of a property",
      Action:   func (c *cli.Context) error {
        key := c.Args().Get(0)
        value := c.Args().Get(1)
        path := c.Args().Get(2)

        p := properties.MustLoadFile(path, properties.UTF8)

        p.MustSet(key, value)

        fo, err := os.Create(path)
        if err != nil {
      		return cli.NewExitError("could not create file", 2)
      	}
        _, e := p.WriteComment(fo, "#", properties.UTF8)
        if e != nil {
      		return cli.NewExitError("could not write file", 3)
      	}
        fo.Close()
        return nil
      },
    },
    {
      Name:     "del",
      Aliases:  []string{"d"},
      Usage:    "deletes a property",
      Action:   func (c *cli.Context) error {
        key := c.Args().Get(0)
        path := c.Args().Get(1)

        p := properties.MustLoadFile(path, properties.UTF8)

        p.Delete(key)

        fo, err := os.Create(path)
        if err != nil {
      		return cli.NewExitError("could not create file", 2)
      	}
        _, e := p.WriteComment(fo, "#", properties.UTF8)
        if e != nil {
      		return cli.NewExitError("could not write file", 3)
      	}
        fo.Close()
        return  nil
      },
    },
  }
  app.Run(os.Args)
}

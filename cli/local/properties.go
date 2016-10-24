package local

import (
	"errors"
	"io"
	"os"

	log "github.com/Sirupsen/logrus"
	"github.com/urfave/cli"
	"golang.org/x/net/context"

	api_operation "github.com/james-nesbitt/wundertools-go/api/operation"
)

// Make CLI flags from operation properties
func CliMakeFlagsFromProperties(props api_operation.Properties) []cli.Flag {
	flags := []cli.Flag{}

	for _, key := range props.Order() {
		prop, _ := props.Get(key)
		value := prop.Get()

		switch value.(type) {
		case string:
			flags = append(flags, cli.StringFlag{
				Name:  prop.Id(),
				Value: value.(string),
				Usage: prop.Description(),
			})
		case []string:
			converted := cli.StringSlice(value.([]string))
			flags = append(flags, cli.StringSliceFlag{
				Name:  prop.Id(),
				Value: &converted,
				Usage: prop.Description(),
			})
		case []byte:
			flags = append(flags, cli.StringFlag{
				Name:  prop.Id(),
				Value: string(value.([]byte)),
				Usage: prop.Description(),
			})
		case int32:
			flags = append(flags, cli.IntFlag{
				Name:  prop.Id(),
				Value: int(value.(int32)),
				Usage: prop.Description(),
			})
		case int64:
			flags = append(flags, cli.Int64Flag{
				Name:  prop.Id(),
				Value: value.(int64),
				Usage: prop.Description(),
			})
		case bool:
			flags = append(flags, cli.BoolFlag{
				Name:  prop.Id(),
				Usage: prop.Description(),
			})
		case io.Writer:
			converted := cli.Generic(&WriterProperty{property: prop})
			flags = append(flags, cli.GenericFlag{
				Name:  prop.Id(),
				Value: converted,
				Usage: prop.Description(),
			})
		case io.Reader:
			converted := cli.Generic(&ReaderProperty{property: prop})
			flags = append(flags, cli.GenericFlag{
				Name:  prop.Id(),
				Value: converted,
				Usage: prop.Description(),
			})
		case context.Context:
			flags = append(flags, cli.IntFlag{
				Name:  prop.Id() + ".timeout",
				Value: 30,
				Usage: "Timeout in seconds. " + prop.Description(),
			})
		default:
			log.WithFields(log.Fields{"id": prop.Id(), "property": prop}).Debug("Unhandled property type for operation")
			converted := cli.Generic(&UnHandledProperty{property: prop})
			flags = append(flags, cli.GenericFlag{
				Name:  prop.Id(),
				Value: converted,
				Usage: prop.Description(),
			})
		}
	}

	return flags
}

// A cli.Generic implementor for un-handled properties
type UnHandledProperty struct {
	property api_operation.Property
}

func (prop *UnHandledProperty) Set(value string) error {
	log.WithFields(log.Fields{"id": prop.property.Id(), "property": prop.property}).Debug("Unhandled property set")
	return nil
}
func (prop *UnHandledProperty) String() string {
	log.WithFields(log.Fields{"id": prop.property.Id(), "property": prop.property}).Debug("Unhandled property retrieve")
	return ""
}

// A cli.Generic implementor for writer properties
type WriterProperty struct {
	property api_operation.Property
}

func (prop *WriterProperty) Set(value string) error {
	switch value {
	case "stdout":
		prop.property.Set(os.Stdout)
	case "stderr":
		prop.property.Set(os.Stdout)
	default:
		return errors.New("Could not interpret flag as an os.Writer")
	}
	return nil
}
func (prop *WriterProperty) String() string {
	log.WithFields(log.Fields{"id": prop.property.Id(), "property": prop.property}).Debug("Unhandled property retrieve")
	return ""
}

// A cli.Generic implementor for reader properties
type ReaderProperty struct {
	property api_operation.Property
}

func (prop *ReaderProperty) Set(value string) error {
	switch value {
	case "stdin":
		prop.property.Set(os.Stdin)
	default:
		return errors.New("Could not interpret flag as an os.Reader")
	}
	return nil
}
func (prop *ReaderProperty) String() string {
	log.WithFields(log.Fields{"id": prop.property.Id(), "property": prop.property}).Debug("Unhandled property retrieve")
	return ""
}

package local

import (
	"io"
	"os"

	log "github.com/Sirupsen/logrus"
	"github.com/urfave/cli"
	"golang.org/x/net/context"

	// libCompose_options "github.com/docker/libcompose/project/options"

	api_operation "github.com/james-nesbitt/wundertools-go/api/operation"
)

// Assign properties from flags back to properties
func CliAssignPropertiesFromFlags(cliContext *cli.Context, props *api_operation.Properties) error {
	for _, key := range props.Order() {
		if !cliContext.IsSet(key) {
			continue
		}

		prop, _ := props.Get(key)

		if prop.Internal() {
			continue
		}

		switch prop.Type() {
		case "string":
			if cliContext.IsSet(key) {
				prop.Set(cliContext.String(key))
			}
		case "[]string":
			if cliContext.IsSet(key) {
				prop.Set(cliContext.StringSlice(key))
			}
		case "[]byte":
			if cliContext.IsSet(key) {
				prop.Set([]byte(cliContext.String(key)))
			}
		case "int":
			if cliContext.IsSet(key) {
				prop.Set(cliContext.Int(key))
			}
		case "int64":
			if cliContext.IsSet(key) {
				prop.Set(cliContext.Int64(key))
			}
		case "bool":
			if cliContext.IsSet(key) {
				prop.Set(cliContext.Bool(key))
			}
		case "io.Writer":
			if cliContext.IsSet(key) {
				switch cliContext.String(key) {
				case "stdout":
					prop.Set(io.Writer(os.Stdout))
				case "stderr":
					prop.Set(io.Writer(os.Stderr))
				}
			}
		case "io.Reader":
			if cliContext.IsSet(key) {
				switch cliContext.String(key) {
				case "stdin":
					prop.Set(io.Reader(os.Stdin))
				}
			}
		case "golang.org/x/net/context.Context":
			if cliContext.IsSet(key + ":duration") {
				duration := cliContext.Duration(key + ".duration")
				if duration > 0 {
					newContext, _ := context.WithTimeout(context.Background(), duration)
					prop.Set(newContext)
				} else {
					prop.Set(context.Background())
				}
			}
		default:
			log.WithFields(log.Fields{"id": prop.Id(), "property": prop, "flag": cliContext.Generic(key)}).Debug("Unhandled property type for operation")
		}
	}

	return nil
}

// Make CLI flags from operation properties
func CliMakeFlagsFromProperties(props api_operation.Properties) []cli.Flag {
	flags := []cli.Flag{}

	for _, key := range props.Order() {
		prop, _ := props.Get(key)

		switch prop.Type() {
		case "string":
			flags = append(flags, cli.StringFlag{
				Name:  prop.Id(),
				Value: prop.Get().(string),
				Usage: prop.Description(),
			})
		case "[]string":
			converted := cli.StringSlice(prop.Get().([]string))
			flags = append(flags, cli.StringSliceFlag{
				Name:  prop.Id(),
				Value: &converted,
				Usage: prop.Description(),
			})
		case "[]byte":
			flags = append(flags, cli.StringFlag{
				Name:  prop.Id(),
				Value: string(prop.Get().([]byte)),
				Usage: prop.Description(),
			})
		case "int32":
			flags = append(flags, cli.IntFlag{
				Name:  prop.Id(),
				Value: int(prop.Get().(int32)),
				Usage: prop.Description(),
			})
		case "int64":
			flags = append(flags, cli.Int64Flag{
				Name:  prop.Id(),
				Value: prop.Get().(int64),
				Usage: prop.Description(),
			})
		case "bool":
			flags = append(flags, cli.BoolFlag{
				Name:  prop.Id(),
				Usage: prop.Description(),
			})
		case "io.Writer":
			flags = append(flags, cli.StringFlag{
				Name:  prop.Id(),
				Value: "",
				Usage: prop.Description(),
			})
		case "io.Reader":
			flags = append(flags, cli.StringFlag{
				Name:  prop.Id(),
				Value: "",
				Usage: prop.Description(),
			})
		case "golang.org/x/net/context.Context":
			flags = append(flags, cli.DurationFlag{
				Name:  prop.Id() + ":duration",
				Usage: "Timeout in seconds. " + prop.Description(),
			})

		default:
			log.WithFields(log.Fields{"id": prop.Id(), "property": prop}).Debug("Unhandled property type for operation")
			converted := cli.Generic(&UnHandledProperty{property: prop})
			flags = append(flags, cli.GenericFlag{
				Name:  prop.Id(),
				Value: converted,
				Usage: "[UNHANDLED] " + prop.Description(),
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

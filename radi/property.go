package radi

import (
	"io"
	"os"

	"context"
	log "github.com/Sirupsen/logrus"
	"gopkg.in/urfave/cli.v2"

	api_operation "github.com/wunderkraut/radi-api/operation"
)

/**
 * Convert between operation.Property and CLI flag formats.
 *
 * There are two functions here to convert CLI flags into
 * operation Property structs (to assign values) and back
 * which allows us to directly use Operations properties as
 * CLI flags without only a thin abstraction layer.
 *
 * The weakness is that we can only abstract property types
 * that we can understand, usually primitives and common
 * types.  If a handler provides a custom property that we
 * don't know about, then we cannot convert it into a flag.
 * This has not caused any problems yet.  These complex cases
 * usually have triggered either a property refactor into
 * simpler primitives, or an alternat approach to setting
 * properties (like relying on config source like yml)
 */

// Assign properties from flags back to properties
func CliAssignPropertiesFromFlags(cliContext *cli.Context, props *api_operation.Properties) error {
	for _, key := range props.Order() {

		if !cliContext.IsSet(key) {
			continue
		}

		prop, _ := props.Get(key)

		// skip any property marked for internal use only
		if !prop.Internal() {

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
			case "context.Context":
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
				log.WithFields(log.Fields{"id": prop.Id(), "property": prop, "flag": cliContext.Generic(key)}).Debug("CLI does not handle any flags for property type")
			}

		}
	}

	return nil
}

// Make CLI flags from operation properties
func CliMakeFlagsFromProperties(props api_operation.Properties) []cli.Flag {
	flags := []cli.Flag{}

	for _, key := range props.Order() {
		prop, _ := props.Get(key)

		// skip any property marked as being for internal use only
		if !prop.Internal() {

			switch prop.Type() {
			case "string":
				flags = append(flags, cli.Flag(&cli.StringFlag{
					Name:  prop.Id(),
					Value: prop.Get().(string),
					Usage: prop.Description(),
				}))
			case "[]string":
				converted := cli.NewStringSlice(prop.Get().([]string)...)
				flags = append(flags, cli.Flag(&cli.StringSliceFlag{
					Name:  prop.Id(),
					Value: converted,
					Usage: prop.Description(),
				}))
			case "[]byte":
				flags = append(flags, cli.Flag(&cli.StringFlag{
					Name:  prop.Id(),
					Value: string(prop.Get().([]byte)),
					Usage: prop.Description(),
				}))
			case "int32":
				flags = append(flags, cli.Flag(&cli.IntFlag{
					Name:  prop.Id(),
					Value: int(prop.Get().(int32)),
					Usage: prop.Description(),
				}))
			case "int64":
				flags = append(flags, cli.Flag(&cli.Int64Flag{
					Name:  prop.Id(),
					Value: prop.Get().(int64),
					Usage: prop.Description(),
				}))
			case "bool":
				flags = append(flags, cli.Flag(&cli.BoolFlag{
					Name:  prop.Id(),
					Usage: prop.Description(),
				}))
			case "io.Writer":
				flags = append(flags, cli.Flag(&cli.StringFlag{
					Name:  prop.Id(),
					Value: "",
					Usage: prop.Description(),
				}))
			case "io.Reader":
				flags = append(flags, cli.Flag(&cli.StringFlag{
					Name:  prop.Id(),
					Value: "",
					Usage: prop.Description(),
				}))
			case "context.Context":
				flags = append(flags, cli.Flag(&cli.DurationFlag{
					Name:  prop.Id() + ":duration",
					Usage: "Timeout in seconds. " + prop.Description(),
				}))

			default:
				log.WithFields(log.Fields{"id": prop.Id(), "property": prop}).Debug("CLI does not yet handle property type for operation")

				/**
				 * originally we were wrapping these unhandled argument types
				 * using a generic handler, but it gave us no real value in
				 * using them, so we stopped.
				 */
				// converted := cli.Generic(&UnHandledProperty{property: prop})
				// flags = append(flags, cli.Flag(&cli.GenericFlag{
				// 	Name:  prop.Id(),
				// 	Value: converted,
				// 	Usage: "[UNHANDLED] " + prop.Description(),
				// }))
			}

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

// Code generated by templ - DO NOT EDIT.

// templ: version: v0.2.778
package dialog

//lint:file-ignore SA4006 This context is only used if a nested component is present.

import "github.com/a-h/templ"
import templruntime "github.com/a-h/templ/runtime"

import (
	"github.com/Oudwins/tailwind-merge-go/pkg/twmerge"
	"github.com/rotemhoresh/shadcn-templ/primitives/dialog"
)

const descriptionClass = "text-sm text-muted-foreground"

type DescriptionProps = dialog.DescriptionProps

func Description(props DescriptionProps) templ.Component {
	props.Class = twmerge.Merge(descriptionClass, props.Class)
	return dialog.Description(props)
}

var _ = templruntime.GeneratedTemplate

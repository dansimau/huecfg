#!/usr/bin/env bash
#
# Generates Go source file for a new subcommand
#

set -eu

for arg in "$@"; do
	varname="$(echo "$arg" | cut -f1 -d '=')"
	varval="$(echo "$arg" | cut -f2 -d '=')"

	eval "$varname='$varval'"
done

#OBJS_NAME=lights
#OBJS_TYPE="[]hue.Light"
#GET_OBJ_FUNC="GetLights()"

OBJS_NAME_FIRST_LETTER=${OBJS_NAME:0:1}
OBJS_NAME_CAMEL=${OBJS_NAME_FIRST_LETTER^^}${OBJS_NAME:1}   # Requires bash 4+

cat <<EOF>${OBJS_NAME}_gen.go
package cmd

//go:generate ./gen_list.sh OBJS_NAME=$OBJS_NAME OBJS_TYPE=$OBJS_TYPE GET_OBJ_FUNC=$GET_OBJ_FUNC

const ${OBJS_NAME}DefaultSortField = "ID"

var (
	${OBJS_NAME}DefaultFields = []string{"ID", "Name"}

	${OBJS_NAME}FieldTransform  fieldTransform
	${OBJS_NAME}HeaderTransform headerTransform
)

func init() {
	_, err := parser.AddCommand("${OBJS_NAME}", "Manage ${OBJS_NAME}", "", &${OBJS_NAME}Cmd{})
	if err != nil {
		panic(err)
	}
}

type ${OBJS_NAME}Cmd struct {
	${OBJS_NAME_CAMEL}List *${OBJS_NAME}ListCmd \`command:"list" alias:"ls" description:"List ${OBJS_NAME}"\`
}
EOF

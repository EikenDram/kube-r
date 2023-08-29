# plumber.R

#' @get /echo
function() {
  Sys.time()
}

#' @plumber
function(pr) {
  pr %>%
    {{- range .Api}}
    pr_mount("{{.Path}}", plumb("/api/{{.File}}")) %>%
    {{- end}}
    pr_mount("/version", plumb("/version.R"))
}
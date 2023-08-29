# plumber.R

#' @get /echo
function() {
  Sys.time()
}

#' @plumber
function(pr) {
  pr %>%
    pr_mount("/test", plumb("/api/test.R")) %>%
    pr_mount("/version", plumb("/version.R"))
}
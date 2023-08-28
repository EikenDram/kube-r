# plumber.R

#' @get /echo
function() {
  Sys.time()
}

#' @plumber
function(pr) {
  pr %>%
    pr_mount("/my_route1", plumb("./file1/myfile.R")) %>%
    pr_mount("/my_route2", plumb("./file1/myfile2.R")) %>%
    pr_mount("/my_route3", plumb("./file2/myfile3.R")) %>%
    pr_mount("/my_route4", plumb("./file2/myfile4.R"))
}
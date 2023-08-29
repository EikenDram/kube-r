# plumber.R

#* Echo back the input
#* @param msg The message to echo
#* @get /info
function() {
  list("KubeR server version 0.0.1", "r-ver version 4.3.1")
}
It's like Imgur, but for sixel images over the command line.

client should work with both .sixels and be able to convert other files for upload
server registers images at randomly generated endpoints made up of lowercase words for easy memorization
server needs rate limiting so that you cant spam image uploads

adapters -> httpserver/grpcserver(?)
cmds -> httpserver/httpclient
domain -> store, serve, convert
specifications -> store, serve, convert

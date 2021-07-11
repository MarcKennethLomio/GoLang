module main

go 1.15

require (
services v1.0.0
models v1.0.0
)

replace services v1.0.0 => ./services
replace models v1.0.0 => ./models
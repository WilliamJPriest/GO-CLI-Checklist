module github.com/WilliamJPriest/checklist

go 1.21.2

replace github.com/WilliamJPriest/checklist/storage => ./storage

require (
	github.com/WilliamJPriest/checklist/cmd v0.0.0-00010101000000-000000000000
	github.com/WilliamJPriest/checklist/storage v0.0.0-00010101000000-000000000000
)

replace github.com/WilliamJPriest/checklist/cmd => ./cmd

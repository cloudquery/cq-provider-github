service          = "github"
output_directory = "."
add_generate     = true

resource "github" "" "users" {
  path = "github.com/google/go-github/v45/github.User"

  column "text_matches" {
    type = "json"
    generate_resolver = true
  }

  options {
    primary_keys = ["id"]
  }

}


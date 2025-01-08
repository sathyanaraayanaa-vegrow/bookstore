json.extract! book, :id, :title, :year, :prize, :genre, :isbn, :created_at, :updated_at
json.url book_url(book, format: :json)

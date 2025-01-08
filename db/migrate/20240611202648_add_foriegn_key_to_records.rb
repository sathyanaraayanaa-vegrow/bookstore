class AddForiegnKeyToRecords < ActiveRecord::Migration[6.0]
  def change
    add_foreign_key :borrowing_records, :books, column: :books_id
    add_foreign_key :borrowing_records, :users, column: :users_id
  end
end

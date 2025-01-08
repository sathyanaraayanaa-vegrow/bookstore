class AddForiegnKeyToRecordsModified < ActiveRecord::Migration[6.0]
  def change
    remove_foreign_key :borrowing_records, column: :books_id
    remove_foreign_key :borrowing_records, column: :users_id
    add_foreign_key :borrowing_records, :books, column: :books_id, on_delete: :cascade
    add_foreign_key :borrowing_records, :users, column: :users_id, on_delete: :cascade
  end
end

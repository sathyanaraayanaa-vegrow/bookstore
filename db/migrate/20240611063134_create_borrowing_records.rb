class CreateBorrowingRecords < ActiveRecord::Migration[6.0]
  def change
    create_table :borrowing_records do |t|
      t.belongs_to :books
      t.belongs_to :users
      t.timestamps
      t.datetime :returnedAt
    end
  end
end

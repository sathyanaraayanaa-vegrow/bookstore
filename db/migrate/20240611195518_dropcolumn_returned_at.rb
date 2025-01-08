class DropcolumnReturnedAt < ActiveRecord::Migration[6.0]
  def change
    remove_column :borrowing_records, :returnedAt
    add_column :borrowing_records, :returned_at, :string
    change_column_default :borrowing_records, :returned_at, nil
  end  
end

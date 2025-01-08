class ChangeColumnReturnedAt < ActiveRecord::Migration[6.0]
  def change
    change_column_default :borrowing_records, :returnedAt, nil
  end
end

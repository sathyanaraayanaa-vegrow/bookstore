class ChangeColumn < ActiveRecord::Migration[6.0]
  def change
    change_column_default :users, :Roles, "user"
  end
end

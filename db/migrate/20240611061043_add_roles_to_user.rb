class AddRolesToUser < ActiveRecord::Migration[6.0]
  def change
    add_column :users, :Roles, :string
  end
end

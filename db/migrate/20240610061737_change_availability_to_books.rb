class ChangeAvailabilityToBooks < ActiveRecord::Migration[6.0]
  def change
    change_column :books, :Availability, :integer
  end
end

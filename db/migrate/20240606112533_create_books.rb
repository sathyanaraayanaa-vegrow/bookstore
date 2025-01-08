class CreateBooks < ActiveRecord::Migration[6.0]
  def change
    create_table :books do |t|
      t.string :title
      t.integer :year
      t.float :prize
      t.string :genre
      t.string :isbn

      t.timestamps
    end
  end
end

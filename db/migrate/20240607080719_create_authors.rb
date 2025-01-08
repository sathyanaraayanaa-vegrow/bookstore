class CreateAuthors < ActiveRecord::Migration[6.0]
  def change
    create_table :authors do |t|
      t.string :name
      t.date :dob
      t.integer :phone

      t.timestamps
    end
  end
end

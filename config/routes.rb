Rails.application.routes.draw do
  devise_for :users, path: '', path_names: {
    sign_in: 'login',
    sign_out: 'logout',
    registration: 'signup',
  },
  controllers: {
    sessions: 'users/sessions',
    registrations: 'users/registrations',
    access: 'users/access'
  }

  # resources :books
  # get 'home/index'
  root 'home#index'
  post '/decode_token', to: 'jwt_auth#decode_token'

  namespace :users do
    get '/', to: 'access#index'
    delete '/:id', to: 'access#destroy', as: 'delete_user'
  end
  # For details on the DSL available within this file, see https://guides.rubyonrails.org/routing.html
end

# path: '', path_names: {
#     sign_in: 'login',
#     sign_out: 'logout',
#     registration: 'signup'
#   },
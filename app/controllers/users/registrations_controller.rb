# frozen_string_literal: true

class Users::RegistrationsController < Devise::RegistrationsController

  respond_to :json

  def create
    role = params[:user][:Roles].to_s.downcase.strip  # Assuming role is passed in params
    
    # Validate role
    unless %w(admin user librarian).include?(role)
      return render json: {
        status: { Msg: "Invalid role. Role must be either 'admin', 'user', or 'librarian'." }
      }, status: :unprocessable_entity
    end

    build_resource(sign_up_params)
    token = request.headers['Authorization'].to_s.split(' ').last

    if %w(admin user librarian).include?(role)
      if %w(user).include?(role) || (token.present? && decode_jwt(token)=="admin")
        resource.Roles = role
        if resource.save
          sign_up(resource_name, resource)
          # @token = Warden::JWTAuth::UserEncoder.new.call(resource, :user, nil)
          @token = request.env['warden-jwt_auth.token']
          decoded_token = JWT.decode(@token, Rails.application.credentials.devise_jwt_secret_key!, true, algorithm: 'HS256').first
          current_user = User.find(decoded_token['sub'])
          decoded_token['role'] = current_user.Roles

          updated_token = JWT.encode(decoded_token, Rails.application.credentials.devise_jwt_secret_key!, 'HS256')

          request.env['warden-jwt_auth.token'] = updated_token

          headers['Authorization'] = updated_token

          UserLogger.create({event: "#{current_user.Roles} Created Successfully.", user_id: current_user.id, email: current_user.email})
          render json: {
            status: { code: 200, Msg: 'Signed up successfully.',
                      token: updated_token,
                      role: current_user.Roles,
                      data: UserSerializer.new(resource).serializable_hash[:data][:attributes] }
          }
        else
          render json: {
            status: { Msg: "User couldn't be created successfully. #{resource.errors.full_messages.to_sentence}" }
          }, status: :unprocessable_entity
        end
      else
        render json: {
          status: { Msg: "Unauthorized Access. #{resource.errors.full_messages.to_sentence}" }
        }, status: :unprocessable_entity
      end
    else
      render json: {
        status: { Msg: "Invalid role. Role must be either 'admin', 'user', or 'librarian'. #{resource.errors.full_messages.to_sentence}" }
      }, status: :unprocessable_entity
    end
  end

  private

    def decode_jwt(token)
      # Implement your JWT decoding logic here
      begin
        decoded = JWT.decode(token, Rails.application.credentials.devise_jwt_secret_key!, true, algorithm: 'HS256').first
        return decoded["role"]  # Assuming the payload is the first element in the decoded array
      rescue JWT::DecodeError
        return nil
      end
    end

  # before_action :configure_sign_up_params, only: [:create]
  # before_action :configure_account_update_params, only: [:update]

  # GET /resource/sign_up
  # def new
  #   super
  # end

  # POST /resource
  # def create
  #   super
  # end

  # GET /resource/edit
  # def edit
  #   super
  # end

  # PUT /resource
  # def update
  #   super
  # end

  # DELETE /resource
  # def destroy
  #   super
  # end

  # GET /resource/cancel
  # Forces the session data which is usually expired after sign
  # in to be expired now. This is useful if the user wants to
  # cancel oauth signing in/up in the middle of the process,
  # removing all OAuth session data.
  # def cancel
  #   super
  # end

  # protected

  # If you have extra params to permit, append them to the sanitizer.
  # def configure_sign_up_params
  #   devise_parameter_sanitizer.permit(:sign_up, keys: [:attribute])
  # end

  # If you have extra params to permit, append them to the sanitizer.
  # def configure_account_update_params
  #   devise_parameter_sanitizer.permit(:account_update, keys: [:attribute])
  # end

  # The path used after sign up.
  # def after_sign_up_path_for(resource)
  #   super(resource)
  # end

  # The path used after sign up for inactive accounts.
  # def after_inactive_sign_up_path_for(resource)
  #   super(resource)
  # end
end

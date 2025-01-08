# frozen_string_literal: true

class Users::SessionsController < Devise::SessionsController
  respond_to :json

  private

  def respond_with(resource, _opt = {})
    @token = request.env['warden-jwt_auth.token']

    decoded_token = JWT.decode(@token, Rails.application.credentials.devise_jwt_secret_key!, true, algorithm: 'HS256').first
    current_user = User.find(decoded_token['sub'])
    decoded_token['role'] = current_user.Roles

    updated_token = JWT.encode(decoded_token, Rails.application.credentials.devise_jwt_secret_key!, 'HS256')

    request.env['warden-jwt_auth.token'] = updated_token

    headers['Authorization'] = updated_token
    UserLogger.create({event: "#{current_user.Roles} Session Created Successfully. Logged In", user_id: current_user.id, email: current_user.email})
    render json: {
      status: {
        code: 200, Msg: 'Logged in successfully.',
        token: updated_token,
        role: current_user.Roles,
        data: {
          user: UserSerializer.new(resource).serializable_hash[:data][:attributes]
        }
      }
    }, status: :ok
  end

  def respond_to_on_destroy
    if request.headers['Authorization'].present?
      jwt_payload = JWT.decode(request.headers['Authorization'].split.last,
                               Rails.application.credentials.devise_jwt_secret_key!).first

      current_user = User.find(jwt_payload['sub'])
    end

    if current_user
      UserLogger.create({event: "#{current_user.Roles} Session Destroyed. Logged Out", user_id: current_user.id, email: current_user.email})
      render json: {
        status: 200,
        Msg: 'Logged out successfully.'
      }, status: :ok
    else
      UserLogger.create({event: "Error: Couldn't find an active session.", user_id: 0, email: ""})
      render json: {
        status: 401,
        Msg: "Error: Couldn't find an active session."
      }, status: :unauthorized
    end
  end


  # before_action :configure_sign_in_params, only: [:create]

  # GET /resource/sign_in
  # def new
  #   super
  # end

  # POST /resource/sign_in
  # def create
  #   super
  # end

  # DELETE /resource/sign_out
  # def destroy
  #   super
  # end

  # protected

  # If you have extra params to permit, append them to the sanitizer.
  # def configure_sign_in_params
  #   devise_parameter_sanitizer.permit(:sign_in, keys: [:attribute])
  # end
end

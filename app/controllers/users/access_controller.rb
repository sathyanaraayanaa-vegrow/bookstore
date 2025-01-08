class Users::AccessController < ApplicationController
    before_action :authorize_request

    def index
        @users = User.where.not(Roles: "admin").order(:Roles).all
        render json: @users
    end

    def destroy
        @user = User.find(params[:id])
        if @user.Roles=="admin"
            render json: { Msg: 'Cannot delete Admin' }, status: :unprocessable_entity
        elsif @user.destroy
          render json: { Msg: 'User deleted successfully' }
        else
          render json: { Msg: 'Failed to delete user' }, status: :unprocessable_entity
        end
    end

    def decode_jwt(token)
        # Implement your JWT decoding logic here
        begin
          decoded = JWT.decode(token, Rails.application.credentials.devise_jwt_secret_key!, true, algorithm: 'HS256').first
          return decoded["role"]  # Assuming the payload is the first element in the decoded array
        rescue JWT::DecodeError
          return nil
        end
      end

    def authorize_request
        token = request.headers['Authorization'].to_s.split(' ').last
        unless token.present? && decode_jwt(token)=="admin"
          render json: { Msg: 'Unauthorized Access' }, status: :unauthorized
        end
      end 

end

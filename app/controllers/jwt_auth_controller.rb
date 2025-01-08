class JwtAuthController < ApplicationController
    before_action :authorize_request

    def decode_token
      token = request.headers['Authorization'].to_s.split(' ').last
      
      if token.blank?
        render json: { Msg: 'Authorization token is required' }, status: :unprocessable_entity
        return
      end
  
      decoded_token = decode_jwt(token)
  
      if decoded_token.nil?
        render json: { Msg: 'Invalid token' }, status: :unauthorized
      else
        render json: { user_id: decoded_token['sub'], role: decoded_token['role'] }
      end
    end
  
    private
  
    def decode_jwt(token)
      # Implement your JWT decoding logic here
      begin
        decoded = JWT.decode(token, Rails.application.credentials.devise_jwt_secret_key!, true, algorithm: 'HS256').first
        return decoded  # Assuming the payload is the first element in the decoded array
      rescue JWT::DecodeError
        return nil
      end
    end

    def authorize_request
        token = request.headers['Authorization'].to_s.split(' ').last
        unless token.present? && decode_jwt(token)
          render json: { Msg: 'Unauthorized' }, status: :unauthorized
        end
      end
  end
  
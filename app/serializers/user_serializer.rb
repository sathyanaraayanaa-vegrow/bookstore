class UserSerializer
  include JSONAPI::Serializer
  attributes :id, :email, :Roles
end

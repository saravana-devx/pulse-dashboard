package auth

type Service struct {
	repo *UserRepository
}

func NewService(repo *UserRepository) *Service {
	return &Service{repo: repo}
}

func (s *Service) CreateUser(req *CreateUserRequest) error {

	exists, err := s.repo.CheckEmailExists(req.Email)

	if err != nil {
		return ErrInternal
	}

	if exists {
		return ErrEmailExists
	}

	if !validateEmail(req.Email) {
		return ErrInvalidEmail
	}

	if err := validatePassword(req.Password); err != nil {
		return wrapWeaKPassword(err)
	}

	hashedPassword, err := HashPassword(req.Password)

	if err != nil {
		return ErrHashingPassword
	}

	user := &UserModel{
		Email:        req.Email,
		PasswordHash: hashedPassword,
	}

	if err := s.repo.CreateUser(user); err != nil {
		return ErrToCreateUser
	}

	/**
	 * TODO : Generate accesstoken/refreshtoken and sent in response
	 */

	return nil
}

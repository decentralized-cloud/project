package service_test

import (
	"context"
	"errors"
	"testing"

	commonErrors "github.com/decentralized-cloud/tenant/common/errors"
	"github.com/decentralized-cloud/tenant/models"
	"github.com/decentralized-cloud/tenant/services/business/contract"
	"github.com/decentralized-cloud/tenant/services/business/service"
	repositoryContract "github.com/decentralized-cloud/tenant/services/repository/contract"
	repsoitoryMock "github.com/decentralized-cloud/tenant/services/repository/mock"
	"github.com/golang/mock/gomock"
	"github.com/lucsky/cuid"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("TenantService Tests", func() {
	var (
		mockCtrl                    *gomock.Controller
		sut                         contract.TenantServiceContract
		mockTenantRepositoryService *repsoitoryMock.MockTenantRepositoryServiceContract
		ctx                         context.Context
	)

	BeforeEach(func() {
		mockCtrl = gomock.NewController(GinkgoT())

		mockTenantRepositoryService = repsoitoryMock.NewMockTenantRepositoryServiceContract(mockCtrl)
		sut, _ = service.NewTenantService(mockTenantRepositoryService)
		ctx = context.Background()
	})

	AfterEach(func() {
		mockCtrl.Finish()
	})

	Context("user tries to instantiate TenantService", func() {
		When("tenant repository service is not provided and NewTenantService is called", func() {
			It("should return ArgumentError", func() {
				service, err := service.NewTenantService(nil)
				Ω(err).Should(HaveOccurred())
				Ω(service).Should(BeNil())

				argumentErr, ok := err.(commonErrors.ArgumentError)
				Ω(ok).Should(BeTrue())
				Ω(argumentErr.ArgumentName).Should(Equal("repositoryService"))
			})
		})

		When("all dependecies are resolved and NewTenantService is called", func() {
			It("should instantiate the new TenantService", func() {
				service, err := service.NewTenantService(mockTenantRepositoryService)
				Ω(err).Should(BeNil())
				Ω(service).ShouldNot(BeNil())
			})
		})
	})

	Describe("CreateTenant", func() {
		var (
			request contract.CreateTenantRequest
		)

		BeforeEach(func() {
			request = contract.CreateTenantRequest{
				Tenant: models.Tenant{
					Name: cuid.New(),
				}}
		})

		Context("tenant service is instantiated", func() {
			When("CreateTenant is called without context", func() {
				It("should return ArgumentError", func() {
					response, err := sut.CreateTenant(nil, &request)
					Ω(err).Should(HaveOccurred())
					Ω(response).Should(BeNil())

					argumentErr, ok := err.(commonErrors.ArgumentError)
					Ω(ok).Should(BeTrue())
					Ω(argumentErr.ArgumentName).Should(Equal("ctx"))
				})
			})

			When("CreateTenant is called without request", func() {
				It("should return ArgumentError", func() {
					response, err := sut.CreateTenant(ctx, nil)
					Ω(err).Should(HaveOccurred())
					Ω(response).Should(BeNil())

					argumentErr, ok := err.(commonErrors.ArgumentError)
					Ω(ok).Should(BeTrue())
					Ω(argumentErr.ArgumentName).Should(Equal("request"))
				})
			})

			When("CreateTenant is called with invalid request", func() {
				It("should return ArgumentError", func() {
					invalidRequest := contract.CreateTenantRequest{
						Tenant: models.Tenant{
							Name: "",
						}}

					response, err := sut.CreateTenant(ctx, &invalidRequest)
					Ω(err).Should(HaveOccurred())
					Ω(response).Should(BeNil())

					argumentErr, ok := err.(commonErrors.ArgumentError)
					Ω(ok).Should(BeTrue())
					Ω(argumentErr.ArgumentName).Should(Equal("request"))
					Ω(argumentErr.ErrorMessage).Should(Equal(invalidRequest.Validate().Error()))
				})
			})

			When("CreateTenant is called with correct input paramters", func() {
				It("should call tenant respository CreateTenant method", func() {
					mockTenantRepositoryService.
						EXPECT().
						CreateTenant(ctx, gomock.Any()).
						Do(func(_ context.Context, mappedRequest *repositoryContract.CreateTenantRequest) {
							Ω(mappedRequest.Tenant).Should(Equal(request.Tenant))
						}).
						Return(&repositoryContract.CreateTenantResponse{TenantID: cuid.New()}, nil)

					_, err := sut.CreateTenant(ctx, &request)
					Ω(err).Should(BeNil())
				})

				When("And tenant repository CreateTenant return TenantAlreadyExistError", func() {
					It("should return TenantAlreadyExistsError", func() {
						mockTenantRepositoryService.
							EXPECT().
							CreateTenant(gomock.Any(), gomock.Any()).
							Return(nil, repositoryContract.NewTenantAlreadyExistsError())

						response, err := sut.CreateTenant(ctx, &request)
						Ω(err).Should(HaveOccurred())
						Ω(response).Should(BeNil())

						_, ok := err.(contract.TenantAlreadyExistsError)
						Ω(ok).Should(BeTrue())
					})
				})

				When("And tenant repository CreateTenant return any error rather than TenantAlreadyExistsError", func() {
					It("should return UnknownError", func() {
						expectedError := errors.New(cuid.New())
						mockTenantRepositoryService.
							EXPECT().
							CreateTenant(gomock.Any(), gomock.Any()).
							Return(nil, expectedError)

						response, err := sut.CreateTenant(ctx, &request)
						Ω(err).Should(HaveOccurred())
						Ω(response).Should(BeNil())

						unknownErr, ok := err.(contract.UnknownError)
						Ω(ok).Should(BeTrue())
						Ω(unknownErr.ErrorMessage).Should(Equal(expectedError.Error()))
					})
				})

				When("And tenant repository CreateTenant return no error", func() {
					It("should return the new tenantID", func() {
						tenantID := cuid.New()
						mockTenantRepositoryService.
							EXPECT().
							CreateTenant(gomock.Any(), gomock.Any()).
							Return(&repositoryContract.CreateTenantResponse{TenantID: tenantID}, nil)

						response, err := sut.CreateTenant(ctx, &request)
						Ω(err).Should(BeNil())
						Ω(response).ShouldNot(BeNil())
						Ω(response.TenantID).Should(Equal(tenantID))
					})
				})
			})
		})
	})

	Describe("ReadTenant", func() {
		var (
			request contract.ReadTenantRequest
		)

		BeforeEach(func() {
			request = contract.ReadTenantRequest{
				TenantID: cuid.New(),
			}
		})

		Context("tenant service is instantiated", func() {
			When("ReadTenant is called without context", func() {
				It("should return ArgumentError", func() {
					response, err := sut.ReadTenant(nil, &request)
					Ω(err).Should(HaveOccurred())
					Ω(response).Should(BeNil())

					argumentErr, ok := err.(commonErrors.ArgumentError)
					Ω(ok).Should(BeTrue())
					Ω(argumentErr.ArgumentName).Should(Equal("ctx"))
				})
			})

			When("ReadTenant is called without request", func() {
				It("should return ArgumentError", func() {
					response, err := sut.ReadTenant(ctx, nil)
					Ω(err).Should(HaveOccurred())
					Ω(response).Should(BeNil())

					argumentErr, ok := err.(commonErrors.ArgumentError)
					Ω(ok).Should(BeTrue())
					Ω(argumentErr.ArgumentName).Should(Equal("request"))
				})
			})

			When("ReadTenant is called with invalid request", func() {
				It("should return ArgumentError", func() {
					invalidRequest := contract.ReadTenantRequest{
						TenantID: "",
					}

					response, err := sut.ReadTenant(ctx, &invalidRequest)
					Ω(err).Should(HaveOccurred())
					Ω(response).Should(BeNil())

					argumentErr, ok := err.(commonErrors.ArgumentError)
					Ω(ok).Should(BeTrue())
					Ω(argumentErr.ArgumentName).Should(Equal("request"))
					Ω(argumentErr.ErrorMessage).Should(Equal(invalidRequest.Validate().Error()))
				})
			})

			When("ReadTenant is called with correct input paramters", func() {
				It("should call tenant respository ReadTenant method", func() {
					mockTenantRepositoryService.
						EXPECT().
						ReadTenant(ctx, gomock.Any()).
						Do(func(_ context.Context, mappedRequest *repositoryContract.ReadTenantRequest) {
							Ω(mappedRequest.TenantID).Should(Equal(request.TenantID))
						}).
						Return(&repositoryContract.ReadTenantResponse{Tenant: models.Tenant{Name: cuid.New()}}, nil)

					_, err := sut.ReadTenant(ctx, &request)
					Ω(err).Should(BeNil())
				})
			})

			When("And tenant repository ReadTenant return TenantNotFoundError", func() {
				It("should return TenantNotFoundError", func() {
					mockTenantRepositoryService.
						EXPECT().
						ReadTenant(gomock.Any(), gomock.Any()).
						Return(nil, repositoryContract.NewTenantNotFoundError(request.TenantID))

					response, err := sut.ReadTenant(ctx, &request)
					Ω(err).Should(HaveOccurred())
					Ω(response).Should(BeNil())

					_, ok := err.(contract.TenantNotFoundError)
					Ω(ok).Should(BeTrue())
				})
			})

			When("And tenant repository ReadTenant return any error rather than TenantNotFoundError", func() {
				It("should return UnknownError", func() {
					expectedError := errors.New(cuid.New())
					mockTenantRepositoryService.
						EXPECT().
						ReadTenant(gomock.Any(), gomock.Any()).
						Return(nil, expectedError)

					response, err := sut.ReadTenant(ctx, &request)
					Ω(err).Should(HaveOccurred())
					Ω(response).Should(BeNil())

					unknownErr, ok := err.(contract.UnknownError)
					Ω(ok).Should(BeTrue())
					Ω(unknownErr.ErrorMessage).Should(Equal(expectedError.Error()))
				})
			})

			When("And tenant repository ReadTenant return no error", func() {
				It("should return the tenantID", func() {
					tenant := models.Tenant{Name: cuid.New()}
					mockTenantRepositoryService.
						EXPECT().
						ReadTenant(gomock.Any(), gomock.Any()).
						Return(&repositoryContract.ReadTenantResponse{Tenant: tenant}, nil)

					response, err := sut.ReadTenant(ctx, &request)
					Ω(err).Should(BeNil())
					Ω(response).ShouldNot(BeNil())
					Ω(response.Tenant.Name).Should(Equal(tenant.Name))
				})
			})
		})
	})

	Describe("UpdateTenant", func() {
		var (
			request contract.UpdateTenantRequest
		)

		BeforeEach(func() {
			request = contract.UpdateTenantRequest{
				TenantID: cuid.New(),
				Tenant:   models.Tenant{Name: cuid.New()},
			}
		})

		Context("tenant service is instantiated", func() {
			When("UpdateTenant is called without context", func() {
				It("should return ArgumentError", func() {
					response, err := sut.UpdateTenant(nil, &request)
					Ω(err).Should(HaveOccurred())
					Ω(response).Should(BeNil())

					argumentErr, ok := err.(commonErrors.ArgumentError)
					Ω(ok).Should(BeTrue())
					Ω(argumentErr.ArgumentName).Should(Equal("ctx"))
				})
			})

			When("UpdateTenant is called without request", func() {
				It("should return ArgumentError", func() {
					response, err := sut.UpdateTenant(ctx, nil)
					Ω(err).Should(HaveOccurred())
					Ω(response).Should(BeNil())

					argumentErr, ok := err.(commonErrors.ArgumentError)
					Ω(ok).Should(BeTrue())
					Ω(argumentErr.ArgumentName).Should(Equal("request"))
				})
			})

			When("UpdateTenant is called with invalid request", func() {
				It("should return ArgumentError", func() {
					invalidRequest := contract.UpdateTenantRequest{
						TenantID: "",
						Tenant:   models.Tenant{Name: ""},
					}

					response, err := sut.UpdateTenant(ctx, &invalidRequest)
					Ω(err).Should(HaveOccurred())
					Ω(response).Should(BeNil())

					argumentErr, ok := err.(commonErrors.ArgumentError)
					Ω(ok).Should(BeTrue())
					Ω(argumentErr.ArgumentName).Should(Equal("request"))
					Ω(argumentErr.ErrorMessage).Should(Equal(invalidRequest.Validate().Error()))
				})
			})

			When("UpdateTenant is called with correct input paramters", func() {
				It("should call tenant respository UpdateTenant method", func() {
					mockTenantRepositoryService.
						EXPECT().
						UpdateTenant(ctx, gomock.Any()).
						Do(func(_ context.Context, mappedRequest *repositoryContract.UpdateTenantRequest) {
							Ω(mappedRequest.TenantID).Should(Equal(request.TenantID))
							Ω(mappedRequest.Tenant.Name).Should(Equal(request.Tenant.Name))
						}).
						Return(&repositoryContract.UpdateTenantResponse{}, nil)

					_, err := sut.UpdateTenant(ctx, &request)
					Ω(err).Should(BeNil())
				})
			})

			When("And tenant repository UpdateTenant return TenantNotFoundError", func() {
				It("should return TenantNotFoundError", func() {
					mockTenantRepositoryService.
						EXPECT().
						UpdateTenant(gomock.Any(), gomock.Any()).
						Return(nil, repositoryContract.NewTenantNotFoundError(request.TenantID))

					response, err := sut.UpdateTenant(ctx, &request)
					Ω(err).Should(HaveOccurred())
					Ω(response).Should(BeNil())

					_, ok := err.(contract.TenantNotFoundError)
					Ω(ok).Should(BeTrue())
				})
			})

			When("And tenant repository UpdateTenant return any error rather than TenantNotFoundError", func() {
				It("should return UnknownError", func() {
					expectedError := errors.New(cuid.New())
					mockTenantRepositoryService.
						EXPECT().
						UpdateTenant(gomock.Any(), gomock.Any()).
						Return(nil, expectedError)

					response, err := sut.UpdateTenant(ctx, &request)
					Ω(err).Should(HaveOccurred())
					Ω(response).Should(BeNil())

					unknownErr, ok := err.(contract.UnknownError)
					Ω(ok).Should(BeTrue())
					Ω(unknownErr.ErrorMessage).Should(Equal(expectedError.Error()))
				})
			})

			When("And tenant repository UpdateTenant return no error", func() {
				It("should return no error", func() {
					mockTenantRepositoryService.
						EXPECT().
						UpdateTenant(gomock.Any(), gomock.Any()).
						Return(&repositoryContract.UpdateTenantResponse{}, nil)

					_, err := sut.UpdateTenant(ctx, &request)
					Ω(err).Should(BeNil())
				})
			})
		})
	})

	Describe("DeleteTenant is called.", func() {
		var (
			request contract.DeleteTenantRequest
		)

		BeforeEach(func() {
			request = contract.DeleteTenantRequest{
				TenantID: cuid.New(),
			}
		})

		Context("tenant service is instantiated", func() {
			When("context is null", func() {
				It("should return ArgumentError and ArgumentName matches the context argument name", func() {
					response, err := sut.DeleteTenant(nil, &request)
					Ω(err).Should(HaveOccurred())
					Ω(response).Should(BeNil())

					argumentErr, ok := err.(commonErrors.ArgumentError)
					Ω(ok).Should(BeTrue())
					Ω(argumentErr.ArgumentName).Should(Equal("ctx"))
				})
			})

			When("request is null", func() {
				It("should return ArgumentError and ArgumentName matches the request argument name", func() {
					response, err := sut.DeleteTenant(ctx, nil)
					Ω(err).Should(HaveOccurred())
					Ω(response).Should(BeNil())

					argumentErr, ok := err.(commonErrors.ArgumentError)
					Ω(ok).Should(BeTrue())
					Ω(argumentErr.ArgumentName).Should(Equal("request"))
				})
			})

			When("request is invalid", func() {
				It("should return ArgumentError and both ArgumentName and ErrorMessage are matched", func() {
					invalidRequest := contract.DeleteTenantRequest{
						TenantID: "",
					}

					response, err := sut.DeleteTenant(ctx, &invalidRequest)
					Ω(err).Should(HaveOccurred())
					Ω(response).Should(BeNil())

					argumentErr, ok := err.(commonErrors.ArgumentError)
					Ω(ok).Should(BeTrue())
					Ω(argumentErr.ArgumentName).Should(Equal("request"))
					Ω(argumentErr.ErrorMessage).Should(Equal(invalidRequest.Validate().Error()))
				})
			})

			When("input paramters are valid", func() {
				It("should call tenant respository DeleteTenant method", func() {
					mockTenantRepositoryService.
						EXPECT().
						DeleteTenant(ctx, gomock.Any()).
						Do(func(_ context.Context, mappedRequest *repositoryContract.DeleteTenantRequest) {
							Ω(mappedRequest.TenantID).Should(Equal(request.TenantID))
						}).
						Return(&repositoryContract.DeleteTenantResponse{}, nil)

					_, err := sut.DeleteTenant(ctx, &request)
					Ω(err).Should(BeNil())
				})
			})

			When("tenant repository DeleteTenant can not find matched tenant", func() {
				It("should return TenantNotFoundError", func() {
					mockTenantRepositoryService.
						EXPECT().
						DeleteTenant(gomock.Any(), gomock.Any()).
						Return(nil, repositoryContract.NewTenantNotFoundError(request.TenantID))

					response, err := sut.DeleteTenant(ctx, &request)
					Ω(err).Should(HaveOccurred())
					Ω(response).Should(BeNil())

					_, ok := err.(contract.TenantNotFoundError)
					Ω(ok).Should(BeTrue())
				})
			})
			When("tenant repository DeleteTenant is faced with any error rather than TenantNotFoundError", func() {
				It("should return UnknownError", func() {
					expectedError := errors.New(cuid.New())
					mockTenantRepositoryService.
						EXPECT().
						DeleteTenant(gomock.Any(), gomock.Any()).
						Return(nil, expectedError)

					response, err := sut.DeleteTenant(ctx, &request)
					Ω(err).Should(HaveOccurred())
					Ω(response).Should(BeNil())

					unknownErr, ok := err.(contract.UnknownError)
					Ω(ok).Should(BeTrue())
					Ω(unknownErr.ErrorMessage).Should(Equal(expectedError.Error()))
				})
			})

			When("tenant repository DeleteTenant compelets successfully", func() {
				It("should return no error", func() {
					mockTenantRepositoryService.
						EXPECT().
						DeleteTenant(gomock.Any(), gomock.Any()).
						Return(&repositoryContract.DeleteTenantResponse{}, nil)

					_, err := sut.DeleteTenant(ctx, &request)
					Ω(err).Should(BeNil())
				})
			})
		})
	})
})

func TestTenantService(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "TenantService Tests")
}

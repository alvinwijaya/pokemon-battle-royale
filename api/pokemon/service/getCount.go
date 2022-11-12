package service

import "github.com/alvinwijaya/pokemon-battle-royale/custom_error"

func (s *pokemonService) GetCount() (uint64, *custom_error.ServiceError) {
	requestResp, serviceErr := s.pokeapiReq.GetAllInPagination()
	if serviceErr != nil {
		return 0, serviceErr
	}

	return requestResp.Count, nil
}

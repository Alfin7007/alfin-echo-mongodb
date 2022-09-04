package bussiness

import "explore/mongodb/features/checklist"

type checklistData struct {
	data checklist.Data
}

func ChecklistBussiness(checkData checklist.Data) checklist.Bussiness {
	return &checklistData{
		data: checkData,
	}
}

func (uc *checklistData) CreateData(core checklist.Core) error {
	err := uc.data.InsertData(core)
	if err != nil {
		return err
	}
	return nil
}

func (uc *checklistData) GetData(useID string) ([]checklist.Core, error) {
	result, err := uc.data.FindData(useID)
	if err != nil {
		return []checklist.Core{}, err
	}
	return result, nil
}

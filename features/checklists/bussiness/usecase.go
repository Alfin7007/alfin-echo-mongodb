package bussiness

import checklist "explore/mongodb/features/checklists"

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

func (uc *checklistData) DeleteChecklist(userID, id string) error {
	err := uc.data.DeleteData(userID, id)
	if err != nil {
		return err
	}
	return nil
}

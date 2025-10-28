package db

import (
	"context"
	"errors"

	"github.com/go-pg/pg/v10"
	"github.com/go-pg/pg/v10/orm"
)

type TgbotLogHubRepo struct {
	db      orm.DB
	filters map[string][]Filter
	sort    map[string][]SortField
	join    map[string][]string
}

// NewTgbotLogHubRepo returns new repository
func NewTgbotLogHubRepo(db orm.DB) TgbotLogHubRepo {
	return TgbotLogHubRepo{
		db:      db,
		filters: map[string][]Filter{},
		sort: map[string][]SortField{
			Tables.AdminRole.Name:     {{Column: Columns.AdminRole.ID, Direction: SortDesc}},
			Tables.Admin.Name:         {{Column: Columns.Admin.ID, Direction: SortDesc}},
			Tables.LogType.Name:       {{Column: Columns.LogType.ID, Direction: SortDesc}},
			Tables.ServiceLog.Name:    {{Column: Columns.ServiceLog.ID, Direction: SortDesc}},
			Tables.ServiceType.Name:   {{Column: Columns.ServiceType.ID, Direction: SortDesc}},
			Tables.ServiceUser.Name:   {{Column: Columns.ServiceUser.ID, Direction: SortDesc}},
			Tables.Service.Name:       {{Column: Columns.Service.ID, Direction: SortDesc}},
			Tables.ServiceAdmins.Name: {{Column: Columns.ServiceAdmins.ServiceID, Direction: SortDesc}},
		},
		join: map[string][]string{
			Tables.AdminRole.Name:     {TableColumns},
			Tables.Admin.Name:         {TableColumns, Columns.Admin.Role},
			Tables.LogType.Name:       {TableColumns},
			Tables.ServiceLog.Name:    {TableColumns, Columns.ServiceLog.Type, Columns.ServiceLog.Service, Columns.ServiceLog.User},
			Tables.ServiceType.Name:   {TableColumns},
			Tables.ServiceUser.Name:   {TableColumns},
			Tables.Service.Name:       {TableColumns, Columns.Service.Type},
			Tables.ServiceAdmins.Name: {TableColumns, Columns.ServiceAdmins.Service, Columns.ServiceAdmins.Admin},
		},
	}
}

// WithTransaction is a function that wraps TgbotLogHubRepo with pg.Tx transaction.
func (tlhr TgbotLogHubRepo) WithTransaction(tx *pg.Tx) TgbotLogHubRepo {
	tlhr.db = tx
	return tlhr
}

// WithEnabledOnly is a function that adds "statusId"=1 as base filter.
func (tlhr TgbotLogHubRepo) WithEnabledOnly() TgbotLogHubRepo {
	f := make(map[string][]Filter, len(tlhr.filters))
	for i := range tlhr.filters {
		f[i] = make([]Filter, len(tlhr.filters[i]))
		copy(f[i], tlhr.filters[i])
		f[i] = append(f[i], StatusEnabledFilter)
	}
	tlhr.filters = f

	return tlhr
}

/*** AdminRole ***/

// FullAdminRole returns full joins with all columns
func (tlhr TgbotLogHubRepo) FullAdminRole() OpFunc {
	return WithColumns(tlhr.join[Tables.AdminRole.Name]...)
}

// DefaultAdminRoleSort returns default sort.
func (tlhr TgbotLogHubRepo) DefaultAdminRoleSort() OpFunc {
	return WithSort(tlhr.sort[Tables.AdminRole.Name]...)
}

// AdminRoleByID is a function that returns AdminRole by ID(s) or nil.
func (tlhr TgbotLogHubRepo) AdminRoleByID(ctx context.Context, id int, ops ...OpFunc) (*AdminRole, error) {
	return tlhr.OneAdminRole(ctx, &AdminRoleSearch{ID: &id}, ops...)
}

// OneAdminRole is a function that returns one AdminRole by filters. It could return pg.ErrMultiRows.
func (tlhr TgbotLogHubRepo) OneAdminRole(ctx context.Context, search *AdminRoleSearch, ops ...OpFunc) (*AdminRole, error) {
	obj := &AdminRole{}
	err := buildQuery(ctx, tlhr.db, obj, search, tlhr.filters[Tables.AdminRole.Name], PagerTwo, ops...).Select()

	if errors.Is(err, pg.ErrMultiRows) {
		return nil, err
	} else if errors.Is(err, pg.ErrNoRows) {
		return nil, nil
	}

	return obj, err
}

// AdminRolesByFilters returns AdminRole list.
func (tlhr TgbotLogHubRepo) AdminRolesByFilters(ctx context.Context, search *AdminRoleSearch, pager Pager, ops ...OpFunc) (adminRoles []AdminRole, err error) {
	err = buildQuery(ctx, tlhr.db, &adminRoles, search, tlhr.filters[Tables.AdminRole.Name], pager, ops...).Select()
	return
}

// CountAdminRoles returns count
func (tlhr TgbotLogHubRepo) CountAdminRoles(ctx context.Context, search *AdminRoleSearch, ops ...OpFunc) (int, error) {
	return buildQuery(ctx, tlhr.db, &AdminRole{}, search, tlhr.filters[Tables.AdminRole.Name], PagerOne, ops...).Count()
}

// AddAdminRole adds AdminRole to DB.
func (tlhr TgbotLogHubRepo) AddAdminRole(ctx context.Context, adminRole *AdminRole, ops ...OpFunc) (*AdminRole, error) {
	q := tlhr.db.ModelContext(ctx, adminRole)
	applyOps(q, ops...)
	_, err := q.Insert()

	return adminRole, err
}

// UpdateAdminRole updates AdminRole in DB.
func (tlhr TgbotLogHubRepo) UpdateAdminRole(ctx context.Context, adminRole *AdminRole, ops ...OpFunc) (bool, error) {
	q := tlhr.db.ModelContext(ctx, adminRole).WherePK()
	if len(ops) == 0 {
		q = q.ExcludeColumn(Columns.AdminRole.ID)
	}
	applyOps(q, ops...)
	res, err := q.Update()
	if err != nil {
		return false, err
	}

	return res.RowsAffected() > 0, err
}

// DeleteAdminRole deletes AdminRole from DB.
func (tlhr TgbotLogHubRepo) DeleteAdminRole(ctx context.Context, id int) (deleted bool, err error) {
	adminRole := &AdminRole{ID: id}

	res, err := tlhr.db.ModelContext(ctx, adminRole).WherePK().Delete()
	if err != nil {
		return false, err
	}

	return res.RowsAffected() > 0, err
}

/*** Admin ***/

// FullAdmin returns full joins with all columns
func (tlhr TgbotLogHubRepo) FullAdmin() OpFunc {
	return WithColumns(tlhr.join[Tables.Admin.Name]...)
}

// DefaultAdminSort returns default sort.
func (tlhr TgbotLogHubRepo) DefaultAdminSort() OpFunc {
	return WithSort(tlhr.sort[Tables.Admin.Name]...)
}

// AdminByID is a function that returns Admin by ID(s) or nil.
func (tlhr TgbotLogHubRepo) AdminByID(ctx context.Context, id int, ops ...OpFunc) (*Admin, error) {
	return tlhr.OneAdmin(ctx, &AdminSearch{ID: &id}, ops...)
}

// OneAdmin is a function that returns one Admin by filters. It could return pg.ErrMultiRows.
func (tlhr TgbotLogHubRepo) OneAdmin(ctx context.Context, search *AdminSearch, ops ...OpFunc) (*Admin, error) {
	obj := &Admin{}
	err := buildQuery(ctx, tlhr.db, obj, search, tlhr.filters[Tables.Admin.Name], PagerTwo, ops...).Select()

	if errors.Is(err, pg.ErrMultiRows) {
		return nil, err
	} else if errors.Is(err, pg.ErrNoRows) {
		return nil, nil
	}

	return obj, err
}

// AdminsByFilters returns Admin list.
func (tlhr TgbotLogHubRepo) AdminsByFilters(ctx context.Context, search *AdminSearch, pager Pager, ops ...OpFunc) (admins []Admin, err error) {
	err = buildQuery(ctx, tlhr.db, &admins, search, tlhr.filters[Tables.Admin.Name], pager, ops...).Select()
	return
}

// CountAdmins returns count
func (tlhr TgbotLogHubRepo) CountAdmins(ctx context.Context, search *AdminSearch, ops ...OpFunc) (int, error) {
	return buildQuery(ctx, tlhr.db, &Admin{}, search, tlhr.filters[Tables.Admin.Name], PagerOne, ops...).Count()
}

// AddAdmin adds Admin to DB.
func (tlhr TgbotLogHubRepo) AddAdmin(ctx context.Context, admin *Admin, ops ...OpFunc) (*Admin, error) {
	q := tlhr.db.ModelContext(ctx, admin)
	applyOps(q, ops...)
	_, err := q.Insert()

	return admin, err
}

// UpdateAdmin updates Admin in DB.
func (tlhr TgbotLogHubRepo) UpdateAdmin(ctx context.Context, admin *Admin, ops ...OpFunc) (bool, error) {
	q := tlhr.db.ModelContext(ctx, admin).WherePK()
	if len(ops) == 0 {
		q = q.ExcludeColumn(Columns.Admin.ID)
	}
	applyOps(q, ops...)
	res, err := q.Update()
	if err != nil {
		return false, err
	}

	return res.RowsAffected() > 0, err
}

// DeleteAdmin deletes Admin from DB.
func (tlhr TgbotLogHubRepo) DeleteAdmin(ctx context.Context, id int) (deleted bool, err error) {
	admin := &Admin{ID: id}

	res, err := tlhr.db.ModelContext(ctx, admin).WherePK().Delete()
	if err != nil {
		return false, err
	}

	return res.RowsAffected() > 0, err
}

/*** LogType ***/

// FullLogType returns full joins with all columns
func (tlhr TgbotLogHubRepo) FullLogType() OpFunc {
	return WithColumns(tlhr.join[Tables.LogType.Name]...)
}

// DefaultLogTypeSort returns default sort.
func (tlhr TgbotLogHubRepo) DefaultLogTypeSort() OpFunc {
	return WithSort(tlhr.sort[Tables.LogType.Name]...)
}

// LogTypeByID is a function that returns LogType by ID(s) or nil.
func (tlhr TgbotLogHubRepo) LogTypeByID(ctx context.Context, id int, ops ...OpFunc) (*LogType, error) {
	return tlhr.OneLogType(ctx, &LogTypeSearch{ID: &id}, ops...)
}

// OneLogType is a function that returns one LogType by filters. It could return pg.ErrMultiRows.
func (tlhr TgbotLogHubRepo) OneLogType(ctx context.Context, search *LogTypeSearch, ops ...OpFunc) (*LogType, error) {
	obj := &LogType{}
	err := buildQuery(ctx, tlhr.db, obj, search, tlhr.filters[Tables.LogType.Name], PagerTwo, ops...).Select()

	if errors.Is(err, pg.ErrMultiRows) {
		return nil, err
	} else if errors.Is(err, pg.ErrNoRows) {
		return nil, nil
	}

	return obj, err
}

// LogTypesByFilters returns LogType list.
func (tlhr TgbotLogHubRepo) LogTypesByFilters(ctx context.Context, search *LogTypeSearch, pager Pager, ops ...OpFunc) (logTypes []LogType, err error) {
	err = buildQuery(ctx, tlhr.db, &logTypes, search, tlhr.filters[Tables.LogType.Name], pager, ops...).Select()
	return
}

// CountLogTypes returns count
func (tlhr TgbotLogHubRepo) CountLogTypes(ctx context.Context, search *LogTypeSearch, ops ...OpFunc) (int, error) {
	return buildQuery(ctx, tlhr.db, &LogType{}, search, tlhr.filters[Tables.LogType.Name], PagerOne, ops...).Count()
}

// AddLogType adds LogType to DB.
func (tlhr TgbotLogHubRepo) AddLogType(ctx context.Context, logType *LogType, ops ...OpFunc) (*LogType, error) {
	q := tlhr.db.ModelContext(ctx, logType)
	applyOps(q, ops...)
	_, err := q.Insert()

	return logType, err
}

// UpdateLogType updates LogType in DB.
func (tlhr TgbotLogHubRepo) UpdateLogType(ctx context.Context, logType *LogType, ops ...OpFunc) (bool, error) {
	q := tlhr.db.ModelContext(ctx, logType).WherePK()
	if len(ops) == 0 {
		q = q.ExcludeColumn(Columns.LogType.ID)
	}
	applyOps(q, ops...)
	res, err := q.Update()
	if err != nil {
		return false, err
	}

	return res.RowsAffected() > 0, err
}

// DeleteLogType deletes LogType from DB.
func (tlhr TgbotLogHubRepo) DeleteLogType(ctx context.Context, id int) (deleted bool, err error) {
	logType := &LogType{ID: id}

	res, err := tlhr.db.ModelContext(ctx, logType).WherePK().Delete()
	if err != nil {
		return false, err
	}

	return res.RowsAffected() > 0, err
}

/*** ServiceLog ***/

// FullServiceLog returns full joins with all columns
func (tlhr TgbotLogHubRepo) FullServiceLog() OpFunc {
	return WithColumns(tlhr.join[Tables.ServiceLog.Name]...)
}

// DefaultServiceLogSort returns default sort.
func (tlhr TgbotLogHubRepo) DefaultServiceLogSort() OpFunc {
	return WithSort(tlhr.sort[Tables.ServiceLog.Name]...)
}

// ServiceLogByID is a function that returns ServiceLog by ID(s) or nil.
func (tlhr TgbotLogHubRepo) ServiceLogByID(ctx context.Context, id int, ops ...OpFunc) (*ServiceLog, error) {
	return tlhr.OneServiceLog(ctx, &ServiceLogSearch{ID: &id}, ops...)
}

// OneServiceLog is a function that returns one ServiceLog by filters. It could return pg.ErrMultiRows.
func (tlhr TgbotLogHubRepo) OneServiceLog(ctx context.Context, search *ServiceLogSearch, ops ...OpFunc) (*ServiceLog, error) {
	obj := &ServiceLog{}
	err := buildQuery(ctx, tlhr.db, obj, search, tlhr.filters[Tables.ServiceLog.Name], PagerTwo, ops...).Select()

	if errors.Is(err, pg.ErrMultiRows) {
		return nil, err
	} else if errors.Is(err, pg.ErrNoRows) {
		return nil, nil
	}

	return obj, err
}

// ServiceLogsByFilters returns ServiceLog list.
func (tlhr TgbotLogHubRepo) ServiceLogsByFilters(ctx context.Context, search *ServiceLogSearch, pager Pager, ops ...OpFunc) (serviceLogs []ServiceLog, err error) {
	err = buildQuery(ctx, tlhr.db, &serviceLogs, search, tlhr.filters[Tables.ServiceLog.Name], pager, ops...).Select()
	return
}

// CountServiceLogs returns count
func (tlhr TgbotLogHubRepo) CountServiceLogs(ctx context.Context, search *ServiceLogSearch, ops ...OpFunc) (int, error) {
	return buildQuery(ctx, tlhr.db, &ServiceLog{}, search, tlhr.filters[Tables.ServiceLog.Name], PagerOne, ops...).Count()
}

// AddServiceLog adds ServiceLog to DB.
func (tlhr TgbotLogHubRepo) AddServiceLog(ctx context.Context, serviceLog *ServiceLog, ops ...OpFunc) (*ServiceLog, error) {
	q := tlhr.db.ModelContext(ctx, serviceLog)
	applyOps(q, ops...)
	_, err := q.Insert()

	return serviceLog, err
}

// UpdateServiceLog updates ServiceLog in DB.
func (tlhr TgbotLogHubRepo) UpdateServiceLog(ctx context.Context, serviceLog *ServiceLog, ops ...OpFunc) (bool, error) {
	q := tlhr.db.ModelContext(ctx, serviceLog).WherePK()
	if len(ops) == 0 {
		q = q.ExcludeColumn(Columns.ServiceLog.ID)
	}
	applyOps(q, ops...)
	res, err := q.Update()
	if err != nil {
		return false, err
	}

	return res.RowsAffected() > 0, err
}

// DeleteServiceLog deletes ServiceLog from DB.
func (tlhr TgbotLogHubRepo) DeleteServiceLog(ctx context.Context, id int) (deleted bool, err error) {
	serviceLog := &ServiceLog{ID: id}

	res, err := tlhr.db.ModelContext(ctx, serviceLog).WherePK().Delete()
	if err != nil {
		return false, err
	}

	return res.RowsAffected() > 0, err
}

/*** ServiceType ***/

// FullServiceType returns full joins with all columns
func (tlhr TgbotLogHubRepo) FullServiceType() OpFunc {
	return WithColumns(tlhr.join[Tables.ServiceType.Name]...)
}

// DefaultServiceTypeSort returns default sort.
func (tlhr TgbotLogHubRepo) DefaultServiceTypeSort() OpFunc {
	return WithSort(tlhr.sort[Tables.ServiceType.Name]...)
}

// ServiceTypeByID is a function that returns ServiceType by ID(s) or nil.
func (tlhr TgbotLogHubRepo) ServiceTypeByID(ctx context.Context, id int, ops ...OpFunc) (*ServiceType, error) {
	return tlhr.OneServiceType(ctx, &ServiceTypeSearch{ID: &id}, ops...)
}

// OneServiceType is a function that returns one ServiceType by filters. It could return pg.ErrMultiRows.
func (tlhr TgbotLogHubRepo) OneServiceType(ctx context.Context, search *ServiceTypeSearch, ops ...OpFunc) (*ServiceType, error) {
	obj := &ServiceType{}
	err := buildQuery(ctx, tlhr.db, obj, search, tlhr.filters[Tables.ServiceType.Name], PagerTwo, ops...).Select()

	if errors.Is(err, pg.ErrMultiRows) {
		return nil, err
	} else if errors.Is(err, pg.ErrNoRows) {
		return nil, nil
	}

	return obj, err
}

// ServiceTypesByFilters returns ServiceType list.
func (tlhr TgbotLogHubRepo) ServiceTypesByFilters(ctx context.Context, search *ServiceTypeSearch, pager Pager, ops ...OpFunc) (serviceTypes []ServiceType, err error) {
	err = buildQuery(ctx, tlhr.db, &serviceTypes, search, tlhr.filters[Tables.ServiceType.Name], pager, ops...).Select()
	return
}

// CountServiceTypes returns count
func (tlhr TgbotLogHubRepo) CountServiceTypes(ctx context.Context, search *ServiceTypeSearch, ops ...OpFunc) (int, error) {
	return buildQuery(ctx, tlhr.db, &ServiceType{}, search, tlhr.filters[Tables.ServiceType.Name], PagerOne, ops...).Count()
}

// AddServiceType adds ServiceType to DB.
func (tlhr TgbotLogHubRepo) AddServiceType(ctx context.Context, serviceType *ServiceType, ops ...OpFunc) (*ServiceType, error) {
	q := tlhr.db.ModelContext(ctx, serviceType)
	applyOps(q, ops...)
	_, err := q.Insert()

	return serviceType, err
}

// UpdateServiceType updates ServiceType in DB.
func (tlhr TgbotLogHubRepo) UpdateServiceType(ctx context.Context, serviceType *ServiceType, ops ...OpFunc) (bool, error) {
	q := tlhr.db.ModelContext(ctx, serviceType).WherePK()
	if len(ops) == 0 {
		q = q.ExcludeColumn(Columns.ServiceType.ID)
	}
	applyOps(q, ops...)
	res, err := q.Update()
	if err != nil {
		return false, err
	}

	return res.RowsAffected() > 0, err
}

// DeleteServiceType deletes ServiceType from DB.
func (tlhr TgbotLogHubRepo) DeleteServiceType(ctx context.Context, id int) (deleted bool, err error) {
	serviceType := &ServiceType{ID: id}

	res, err := tlhr.db.ModelContext(ctx, serviceType).WherePK().Delete()
	if err != nil {
		return false, err
	}

	return res.RowsAffected() > 0, err
}

/*** ServiceUser ***/

// FullServiceUser returns full joins with all columns
func (tlhr TgbotLogHubRepo) FullServiceUser() OpFunc {
	return WithColumns(tlhr.join[Tables.ServiceUser.Name]...)
}

// DefaultServiceUserSort returns default sort.
func (tlhr TgbotLogHubRepo) DefaultServiceUserSort() OpFunc {
	return WithSort(tlhr.sort[Tables.ServiceUser.Name]...)
}

// ServiceUserByID is a function that returns ServiceUser by ID(s) or nil.
func (tlhr TgbotLogHubRepo) ServiceUserByID(ctx context.Context, id int, ops ...OpFunc) (*ServiceUser, error) {
	return tlhr.OneServiceUser(ctx, &ServiceUserSearch{ID: &id}, ops...)
}

// OneServiceUser is a function that returns one ServiceUser by filters. It could return pg.ErrMultiRows.
func (tlhr TgbotLogHubRepo) OneServiceUser(ctx context.Context, search *ServiceUserSearch, ops ...OpFunc) (*ServiceUser, error) {
	obj := &ServiceUser{}
	err := buildQuery(ctx, tlhr.db, obj, search, tlhr.filters[Tables.ServiceUser.Name], PagerTwo, ops...).Select()

	if errors.Is(err, pg.ErrMultiRows) {
		return nil, err
	} else if errors.Is(err, pg.ErrNoRows) {
		return nil, nil
	}

	return obj, err
}

// ServiceUsersByFilters returns ServiceUser list.
func (tlhr TgbotLogHubRepo) ServiceUsersByFilters(ctx context.Context, search *ServiceUserSearch, pager Pager, ops ...OpFunc) (serviceUsers []ServiceUser, err error) {
	err = buildQuery(ctx, tlhr.db, &serviceUsers, search, tlhr.filters[Tables.ServiceUser.Name], pager, ops...).Select()
	return
}

// CountServiceUsers returns count
func (tlhr TgbotLogHubRepo) CountServiceUsers(ctx context.Context, search *ServiceUserSearch, ops ...OpFunc) (int, error) {
	return buildQuery(ctx, tlhr.db, &ServiceUser{}, search, tlhr.filters[Tables.ServiceUser.Name], PagerOne, ops...).Count()
}

// AddServiceUser adds ServiceUser to DB.
func (tlhr TgbotLogHubRepo) AddServiceUser(ctx context.Context, serviceUser *ServiceUser, ops ...OpFunc) (*ServiceUser, error) {
	q := tlhr.db.ModelContext(ctx, serviceUser)
	applyOps(q, ops...)
	_, err := q.Insert()

	return serviceUser, err
}

// UpdateServiceUser updates ServiceUser in DB.
func (tlhr TgbotLogHubRepo) UpdateServiceUser(ctx context.Context, serviceUser *ServiceUser, ops ...OpFunc) (bool, error) {
	q := tlhr.db.ModelContext(ctx, serviceUser).WherePK()
	if len(ops) == 0 {
		q = q.ExcludeColumn(Columns.ServiceUser.ID)
	}
	applyOps(q, ops...)
	res, err := q.Update()
	if err != nil {
		return false, err
	}

	return res.RowsAffected() > 0, err
}

// DeleteServiceUser deletes ServiceUser from DB.
func (tlhr TgbotLogHubRepo) DeleteServiceUser(ctx context.Context, id int) (deleted bool, err error) {
	serviceUser := &ServiceUser{ID: id}

	res, err := tlhr.db.ModelContext(ctx, serviceUser).WherePK().Delete()
	if err != nil {
		return false, err
	}

	return res.RowsAffected() > 0, err
}

/*** Service ***/

// FullService returns full joins with all columns
func (tlhr TgbotLogHubRepo) FullService() OpFunc {
	return WithColumns(tlhr.join[Tables.Service.Name]...)
}

// DefaultServiceSort returns default sort.
func (tlhr TgbotLogHubRepo) DefaultServiceSort() OpFunc {
	return WithSort(tlhr.sort[Tables.Service.Name]...)
}

// ServiceByID is a function that returns Service by ID(s) or nil.
func (tlhr TgbotLogHubRepo) ServiceByID(ctx context.Context, id int, ops ...OpFunc) (*Service, error) {
	return tlhr.OneService(ctx, &ServiceSearch{ID: &id}, ops...)
}

// OneService is a function that returns one Service by filters. It could return pg.ErrMultiRows.
func (tlhr TgbotLogHubRepo) OneService(ctx context.Context, search *ServiceSearch, ops ...OpFunc) (*Service, error) {
	obj := &Service{}
	err := buildQuery(ctx, tlhr.db, obj, search, tlhr.filters[Tables.Service.Name], PagerTwo, ops...).Select()

	if errors.Is(err, pg.ErrMultiRows) {
		return nil, err
	} else if errors.Is(err, pg.ErrNoRows) {
		return nil, nil
	}

	return obj, err
}

// ServicesByFilters returns Service list.
func (tlhr TgbotLogHubRepo) ServicesByFilters(ctx context.Context, search *ServiceSearch, pager Pager, ops ...OpFunc) (services []Service, err error) {
	err = buildQuery(ctx, tlhr.db, &services, search, tlhr.filters[Tables.Service.Name], pager, ops...).Select()
	return
}

// CountServices returns count
func (tlhr TgbotLogHubRepo) CountServices(ctx context.Context, search *ServiceSearch, ops ...OpFunc) (int, error) {
	return buildQuery(ctx, tlhr.db, &Service{}, search, tlhr.filters[Tables.Service.Name], PagerOne, ops...).Count()
}

// AddService adds Service to DB.
func (tlhr TgbotLogHubRepo) AddService(ctx context.Context, service *Service, ops ...OpFunc) (*Service, error) {
	q := tlhr.db.ModelContext(ctx, service)
	applyOps(q, ops...)
	_, err := q.Insert()

	return service, err
}

// UpdateService updates Service in DB.
func (tlhr TgbotLogHubRepo) UpdateService(ctx context.Context, service *Service, ops ...OpFunc) (bool, error) {
	q := tlhr.db.ModelContext(ctx, service).WherePK()
	if len(ops) == 0 {
		q = q.ExcludeColumn(Columns.Service.ID)
	}
	applyOps(q, ops...)
	res, err := q.Update()
	if err != nil {
		return false, err
	}

	return res.RowsAffected() > 0, err
}

// DeleteService deletes Service from DB.
func (tlhr TgbotLogHubRepo) DeleteService(ctx context.Context, id int) (deleted bool, err error) {
	service := &Service{ID: id}

	res, err := tlhr.db.ModelContext(ctx, service).WherePK().Delete()
	if err != nil {
		return false, err
	}

	return res.RowsAffected() > 0, err
}

/*** ServiceAdmins ***/

// FullServiceAdmins returns full joins with all columns
func (tlhr TgbotLogHubRepo) FullServiceAdmins() OpFunc {
	return WithColumns(tlhr.join[Tables.ServiceAdmins.Name]...)
}

// DefaultServiceAdminsSort returns default sort.
func (tlhr TgbotLogHubRepo) DefaultServiceAdminsSort() OpFunc {
	return WithSort(tlhr.sort[Tables.ServiceAdmins.Name]...)
}

// ServiceAdminsByID is a function that returns ServiceAdmins by ID(s) or nil.
func (tlhr TgbotLogHubRepo) ServiceAdminsByID(ctx context.Context, serviceID int, adminID int, ops ...OpFunc) (*ServiceAdmins, error) {
	return tlhr.OneServiceAdmins(ctx, &ServiceAdminsSearch{ServiceID: &serviceID, AdminID: &adminID}, ops...)
}

// OneServiceAdmins is a function that returns one ServiceAdmins by filters. It could return pg.ErrMultiRows.
func (tlhr TgbotLogHubRepo) OneServiceAdmins(ctx context.Context, search *ServiceAdminsSearch, ops ...OpFunc) (*ServiceAdmins, error) {
	obj := &ServiceAdmins{}
	err := buildQuery(ctx, tlhr.db, obj, search, tlhr.filters[Tables.ServiceAdmins.Name], PagerTwo, ops...).Select()

	if errors.Is(err, pg.ErrMultiRows) {
		return nil, err
	} else if errors.Is(err, pg.ErrNoRows) {
		return nil, nil
	}

	return obj, err
}

// ServiceAdminsByFilters returns ServiceAdmins list.
func (tlhr TgbotLogHubRepo) ServiceAdminsByFilters(ctx context.Context, search *ServiceAdminsSearch, pager Pager, ops ...OpFunc) (serviceAdminsList []ServiceAdmins, err error) {
	err = buildQuery(ctx, tlhr.db, &serviceAdminsList, search, tlhr.filters[Tables.ServiceAdmins.Name], pager, ops...).Select()
	return
}

// CountServiceAdmins returns count
func (tlhr TgbotLogHubRepo) CountServiceAdmins(ctx context.Context, search *ServiceAdminsSearch, ops ...OpFunc) (int, error) {
	return buildQuery(ctx, tlhr.db, &ServiceAdmins{}, search, tlhr.filters[Tables.ServiceAdmins.Name], PagerOne, ops...).Count()
}

// AddServiceAdmins adds ServiceAdmins to DB.
func (tlhr TgbotLogHubRepo) AddServiceAdmins(ctx context.Context, serviceAdmins *ServiceAdmins, ops ...OpFunc) (*ServiceAdmins, error) {
	q := tlhr.db.ModelContext(ctx, serviceAdmins)
	applyOps(q, ops...)
	_, err := q.Insert()

	return serviceAdmins, err
}

// UpdateServiceAdmins updates ServiceAdmins in DB.
func (tlhr TgbotLogHubRepo) UpdateServiceAdmins(ctx context.Context, serviceAdmins *ServiceAdmins, ops ...OpFunc) (bool, error) {
	q := tlhr.db.ModelContext(ctx, serviceAdmins).WherePK()
	if len(ops) == 0 {
		q = q.ExcludeColumn(Columns.ServiceAdmins.ServiceID, Columns.ServiceAdmins.AdminID)
	}
	applyOps(q, ops...)
	res, err := q.Update()
	if err != nil {
		return false, err
	}

	return res.RowsAffected() > 0, err
}

// DeleteServiceAdmins deletes ServiceAdmins from DB.
func (tlhr TgbotLogHubRepo) DeleteServiceAdmins(ctx context.Context, serviceID int, adminID int) (deleted bool, err error) {
	serviceAdmins := &ServiceAdmins{ServiceID: serviceID, AdminID: adminID}

	res, err := tlhr.db.ModelContext(ctx, serviceAdmins).WherePK().Delete()
	if err != nil {
		return false, err
	}

	return res.RowsAffected() > 0, err
}

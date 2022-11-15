package todoService

import (
	"acc/app/model/bean"
	"acc/app/model/dao"
	"acc/app/model/dto"
	"acc/app/model/entity"
	"acc/lib/errCode"
	"acc/lib/orm"
	"acc/lib/response"
	"time"
)

func TodoAdd(data *dto.TodoReq) (res interface{}, biz *errCode.ApiError) {
	q := dao.Use(orm.DB())
	//开启事务
	err := q.Transaction(func(tx *dao.Query) error {
		qtodo := tx.Todo
		qtag := tx.Tag
		qlink := tx.TodoTag
		//先查uid是否存在

		//插入 todo
		err := qtodo.Create(&data.Todo)
		if err != nil {
			biz = errCode.NewErrorMsg("操作失败")
			return err
		}

		//插入 link 和 tag
		for i, tag := range data.Tags {
			//插入 tag,存在则返回
			tag, err := qtag.Where(qtag.Name.Eq(tag.Name)).FirstOrCreate()
			if err != nil {
				biz = errCode.NewErrorMsg("创建tag错误")
				return err
			}

			tt := entity.TodoTag{TodoID: data.Todo.ID, TagID: tag.ID}
			err = qlink.Create(&tt)
			if err != nil {
				biz = errCode.NewErrorMsg("创建link错误")
				return err
			}
			data.Tags[i] = tag
		}

		//返回 nil 提交事务,返回任何错误都会回滚事务
		return nil
	})

	if err != nil {
		return
	}

	res = data
	return
}

func TodoDelete(id int) (res interface{}, biz *errCode.ApiError) {
	q := dao.Use(orm.DB())
	//开启事务
	err := q.Transaction(func(tx *dao.Query) error {
		qtodo := tx.Todo
		qlink := tx.TodoTag

		data, err := qtodo.Where(qtodo.ID.Eq(id)).First()
		if err != nil {
			biz = errCode.OnError(err)
			return err
		}
		//先查link
		links, err := qlink.Where(qlink.TodoID.Eq(data.ID)).Find()

		//删除 todo
		info, err := qtodo.Delete(data)
		if err != nil {
			biz = errCode.OnError(err)
			return err
		}

		//删除link,不删tag
		for _, l := range links {
			_, err := qlink.Delete(l)
			if err != nil {
				return err
			}
		}

		res = info
		//返回 nil 提交事务,返回任何错误都会回滚事务
		return nil
	})

	if err != nil {
		return
	}

	return
}

func TodoUpdate(data *dto.TodoReq) (res interface{}, biz *errCode.ApiError) {
	q := dao.Use(orm.DB())
	//开启事务
	err := q.Transaction(func(tx *dao.Query) error {
		qtodo := tx.Todo
		qtag := tx.Tag
		qlink := tx.TodoTag

		//先查link
		links, err := qlink.Where(qlink.TodoID.Eq(data.Todo.ID)).Find()

		//更新 todo
		err = qtodo.Save(&data.Todo)
		if err != nil {
			biz = errCode.NewErrorMsg("操作失败")
			return err
		}

		//删除link,不删tag
		for _, l := range links {
			_, err := qlink.Delete(l)
			if err != nil {
				return err
			}
		}
		//插入 link
		for i, tag := range data.Tags {
			//插入 tag,存在则返回
			tag, err := qtag.Where(qtag.Name.Eq(tag.Name)).FirstOrCreate()
			if err != nil {
				biz = errCode.NewErrorMsg("创建tag错误")
				return err
			}

			tt := entity.TodoTag{TodoID: data.Todo.ID, TagID: tag.ID}
			err = qlink.Create(&tt)
			if err != nil {
				biz = errCode.NewErrorMsg("创建link错误")
				return err
			}
			data.Tags[i] = tag
		}

		//返回 nil 提交事务,返回任何错误都会回滚事务
		return nil
	})

	if err != nil {
		biz = errCode.NewErrorMsg("操作失败")
		return
	}

	res = data
	return
}

func TodoQuery(page int, condition dto.Condition) (result *response.PageResult, err error) {
	qt := dao.Use(orm.DB()).Todo
	qu := dao.Use(orm.DB()).UserInfo

	var qw = qt.Where()
	if condition.Keywords != "" {
		qw = qw.Where(qt.Title.Like("%" + condition.Keywords + "%"))
	}

	if condition.StartTime != "" {
		start, _ := time.Parse(time.RFC3339, condition.StartTime)
		qw = qw.Where(qt.StartTime.Gte(start))
	}

	if condition.EndTime != "" {
		end, _ := time.Parse(time.RFC3339, condition.EndTime)
		qw = qw.Where(qt.StartTime.Lte(end))
	}

	limit := 5
	offset := page*limit - limit

	//查todo，限制用户只能查询自己的todo
	todos, total, err := qw.Where(qt.Where(qt.UID.Eq(condition.Uid))).FindByPage(offset, limit)
	if err != nil {
		return nil, err
	}
	//查用户
	userinfo, err := qu.Where(qu.ID.Eq(condition.Uid)).First()
	if err != nil {
		return nil, err
	}

	result = &response.PageResult{
		Datas: convertTodoBean(userinfo, todos),
		Size:  limit,
		Page:  page,
		Total: total,
	}

	return result, nil
}

func convertTodoBean(user *entity.UserInfo, list []*entity.Todo) []bean.TodoInfoResp {
	var todoinfos []bean.TodoInfoResp
	for _, todo := range list {
		qtag := dao.Use(orm.DB()).Tag
		qlink := dao.Use(orm.DB()).TodoTag
		//先查link
		links, err := qlink.Where(qlink.TodoID.Eq(todo.ID)).Find()
		if err != nil {
			return nil
		}

		var ids []int
		for _, link := range links {
			//查tag
			ids = append(ids, link.TagID)
		}

		tags, err := qtag.Where(qtag.ID.In(ids...)).Find()

		todoinfos = append(todoinfos, bean.TodoInfoResp{
			Username: user.Nickname,
			Tags:     tags,
			Details:  todo,
		})
	}
	return todoinfos
}

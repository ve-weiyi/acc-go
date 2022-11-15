import {request} from "@/utils/service"

interface ITodoDataApi {
  todo?: TodoBean
  tag?: TagBean[]
}


interface TodoBean {
  id: number
  uid: number
  priority: number
  done: number
  title: string
  content?: string
  start_time?: string
}

interface TagBean {
  id: number
  name?: string
}

interface IGetTableDataApi {
  // /** 当前页码 */
  // currentPage: number
  // /** 查询条数 */
  // size: number
  id?:number
  /** 查询参数 */
  keywords?: string
  startTime?: string
  endTime?: string
}

/** 增 */
export function createTableDataApi(data: ITodoDataApi) {
  return request({
    url: "/todo/add",
    method: "post",
    data
  })
}

/** 删 */
export function deleteTableDataApi(params: IGetTableDataApi) {
  return request({
    url: `/todo/delete`,
    method: "delete",
    params
  })
}

/** 改 */
export function updateTableDataApi(data: ITodoDataApi) {
  return request({
    url: "/todo/update",
    method: "post",
    data
  })
}

/** 查 */
export function getTableDataApi(page: number, params: IGetTableDataApi) {
  return request({
    url: "/todo/list/" + page,
    method: "get",
    params
  })
}

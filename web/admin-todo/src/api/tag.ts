import { request } from "@/utils/service"

interface TagBean {
  id: number
  name ?: string
}


interface ICondition {
  // /** 当前页码 */
  // currentPage: number
  // /** 查询条数 */
  // size: number
  /** 查询参数 */
  keywords?: string
}

/** 增 */
export function createTagApi(data: TagBean) {
  return request({
    url: "/tag/add",
    method: "post",
    data
  })
}

/** 删 */
export function deleteTagApi(data: TagBean) {
  return request({
    url: "/tag/delete",
    method: "delete",
    data
  })
}

/** 改 */
export function updateTagApi(data: TagBean) {
  return request({
    url: "/tag/update",
    method: "post",
    data
  })
}

/** 查 */
export function getTagListApi(page :number,params: ICondition) {
  return request({
    url: "/tag/list/"+page,
    method: "get",
    params
  })
}

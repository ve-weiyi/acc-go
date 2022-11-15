export interface IListItem {
  avatar?: string
  title: string
  datetime?: string
  description?: string
  status?: "" | "success" | "info" | "warning" | "danger"
  extra?: string
}

export const notifyData: IListItem[] = [
  {
    avatar: "https://gw.alipayobjects.com/zos/rmsportal/OKJXDXrmkNshAMvwtvhu.png",
    keywords: "V3 Admin Vite 上线啦",
    datetime: "半年前",
    description:
      "一个免费开源的中后台管理系统基础解决方案，基于 Vue3、TypeScript、Element Plus、Pinia 和 Vite 等主流技术"
  },
  {
    avatar: "https://gw.alipayobjects.com/zos/rmsportal/OKJXDXrmkNshAMvwtvhu.png",
    keywords: "V3 Admin 上线啦",
    datetime: "一年前",
    description: "一个中后台管理系统基础解决方案，基于 Vue3、TypeScript、Element Plus 和 Pinia"
  }
]

export const messageData: IListItem[] = [
  {
    avatar: "https://gw.alipayobjects.com/zos/rmsportal/ThXAXghbEsBCCSDihZxY.png",
    keywords: "来自楚门的世界",
    description: "如果再也不能见到你，祝你早安、午安和晚安",
    datetime: "1998-06-05"
  },
  {
    avatar: "https://gw.alipayobjects.com/zos/rmsportal/ThXAXghbEsBCCSDihZxY.png",
    keywords: "来自大话西游",
    description: "如果非要在这份爱上加上一个期限，我希望是一万年",
    datetime: "1995-02-04"
  },
  {
    avatar: "https://gw.alipayobjects.com/zos/rmsportal/ThXAXghbEsBCCSDihZxY.png",
    keywords: "来自龙猫",
    description: "心存善意，定能途遇天使",
    datetime: "1988-04-16"
  }
]

export const todoData: IListItem[] = [
  {
    keywords: "任务名称",
    description: "这家伙很懒，什么都没留下",
    extra: "未开始",
    status: "info"
  },
  {
    keywords: "任务名称",
    description: "这家伙很懒，什么都没留下",
    extra: "进行中",
    status: ""
  },
  {
    keywords: "任务名称",
    description: "这家伙很懒，什么都没留下",
    extra: "已超时",
    status: "danger"
  }
]

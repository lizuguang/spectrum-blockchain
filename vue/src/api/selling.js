import request from '@/utils/request'

// 查询出租(可查询所有，也可根据发起出租人查询)(发起的)
export function querySellingList(data) {
  return request({
    url: '/querySellingList',
    method: 'post',
    data
  })
}

// 根据参与出租人、买家(买家AccountId)查询出租(参与的)
export function querySellingListByBuyer(data) {
  return request({
    url: '/querySellingListByBuyer',
    method: 'post',
    data
  })
}

// 买家购买
export function createSellingByBuy(data) {
  return request({
    url: '/createSellingByBuy',
    method: 'post',
    data
  })
}

// 更新出租状态（买家确认、买卖家取消）Status取值为 完成"done"、取消"cancelled" 当处于出租中状态，卖家要取消时，buyer为""空
export function updateSelling(data) {
  return request({
    url: '/updateSelling',
    method: 'post',
    data
  })
}

// 发起出租
export function createSelling(data) {
  return request({
    url: '/createSelling',
    method: 'post',
    data
  })
}

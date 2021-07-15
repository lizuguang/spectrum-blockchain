import request from '@/utils/request'

// 查询转让列表(可查询所有，也可根据发起转让人查询)
export function queryDonatingList(data) {
  return request({
    url: '/queryDonatingList',
    method: 'post',
    data
  })
}

// 根据接收人(接收人AccountId)查询转让(接收的)(供接收人查询)
export function queryDonatingListByGrantee(data) {
  return request({
    url: '/queryDonatingListByGrantee',
    method: 'post',
    data
  })
}

// 更新转让状态（确认接收、取消） Status取值为 完成"done"、取消"cancelled"
export function updateDonating(data) {
  return request({
    url: '/updateDonating',
    method: 'post',
    data
  })
}

// 发起转让
export function createDonating(data) {
  return request({
    url: '/createDonating',
    method: 'post',
    data
  })
}

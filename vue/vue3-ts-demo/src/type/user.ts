export interface ListInt {
    id: number,
    nickName: string,
    userName: string
    role: RoleInt[]
}

export interface RoleInt {
    role: number,
    roleName: string
}
export interface selectData {
    nickName: string
    role: number
}

export interface RoleListInt {
    authority:number
    roleId:number
    roleName:string
}
export interface ActiveInt{
    id: number,
    nickName: string,
    userName: string
    role: number[]
}
export class InitData {
    list: ListInt[] = []
    selectData: selectData = { nickName: "", role: 0 }
    roleList:RoleListInt[]=[]
    isShow:boolean=false
    active:ActiveInt={
        id: 0,
        nickName: "",
        userName: "",
        role: []
    }
}
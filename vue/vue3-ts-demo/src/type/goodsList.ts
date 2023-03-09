export interface ListInt {
    userId: number,
    id: number,
    title: string,
    introduce: string,
}
export interface selectData {
    title:string,
    introduce: string,
    page:number,
    count:number,
    pagesize:number
}
export class InitData {
    selectData: selectData = {
        title:"",
        introduce:"",
        page:1,
        count:0,
        pagesize:5
    }
    list: ListInt[] = []
}
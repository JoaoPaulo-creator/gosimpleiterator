```ts
type ProductObj = {
        id: string
        name: string
        price: number
}

let list: ProductObj[] = [
        { 'id': 'alksdj', 'name': 'ab1', 'price': 12.46 },
        { 'id': 'lkjsop', 'name': 'ab2', 'price': 387.22 },
        { 'id': 'alksjd', 'name': 'ab3', 'price': 98.37 },
        { 'id': 'llksdj', 'name': 'ab4', 'price': 12353 },
]

let mappedList = list.map((item: ProductObj) => item.id)
console.log(mappedList)

type ProductObj = {
        name: string
        address: Address[]
}


type Address = {
        streetName: string
        streetNumber: string
        neighbourhood: string
        addressType?: string
}


let list: ProductObj = {
        'name': 'Zé', address: [
                { streetName: "Marechal Deorodo", streetNumber: '630', neighbourhood: 'Centro' },
                { streetName: "Avenida Mariano Peixoto", streetNumber: '10', neighbourhood: 'Centro' },
                { streetName: "Av. Sete de Setembro", streetNumber: '80', neighbourhood: 'Batel' },
                { streetName: "Av. Republica Argentina", streetNumber: '17g', neighbourhood: 'Agua verde' },
        ]
}


/**
(method) Array<Address>.map<string>(callbackfn: (value: Address, index: number, array: Address[]) => string, thisArg?: any): string[]

Calls a defined callback function on each element of an array, and returns an array that contains the results.

*@param* `callbackfn` — A function that accepts up to three arguments. The map method calls the callbackfn function one time for each element in the array.  
*@param* `thisArg` — An object to which the this keyword can refer in the callbackfn function. If thisArg is omitted, undefined is used as the this value.
*/

let mappedList = list.address.map((item: Address) => {
        return {
                rua: item.streetName,
                bairro: item.neighbourhood
        }
})


let filteredList = list.address.filter((item: Address) => item.addressType == 'trabalho')

console.log(mappedList)



```

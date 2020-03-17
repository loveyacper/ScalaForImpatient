package ch1stackqueue

type Pet interface {
    GetPetType() string
}

type Cat struct {
}

func (c Cat) GetPetType() string {
    return "cat"
}

type Dog struct {
}

func (d Dog) GetPetType() string {
    return "dog"
}

type catItem struct {
    seq int64
    cat Cat
}
type dogItem struct {
    seq int64
    dog Dog
}

// 队列维护一个seq，单调递增，记录每个pet入队seq
// 维护两个子队列，分别保存dog和cat
// pollAll则是比较两个子队列的头部，谁的seq小，就弹出谁
type CatDogQueue struct {
    dogs []dogItem
    cats []catItem
    seq int64
}

func (q *CatDogQueue) Add(pet Pet) {
    if pet.GetPetType() == "cat" {
        seq := q.seq
        item := catItem{seq : seq, cat : pet.(Cat)}
        q.cats = append(q.cats, item)
        q.seq++
    } else if pet.GetPetType() == "dog" {
        seq := q.seq
        item := dogItem{seq : seq, dog : pet.(Dog)}
        q.dogs = append(q.dogs, item)
        q.seq++
    } else {
        panic("Wrong pet type")
    }
}

func (q *CatDogQueue) IsDogEmpty() bool {
    return len(q.dogs) == 0
}

func (q *CatDogQueue) IsCatEmpty() bool {
    return len(q.cats) == 0
}

func (q *CatDogQueue) IsEmpty() bool {
    return q.IsDogEmpty() && q.IsCatEmpty()
}

func (q *CatDogQueue) PollCat() Cat {
    if q.IsCatEmpty() {
        panic("No cats")
    }

    cat := q.cats[0].cat
    q.cats = q.cats[1:]

    return cat
}

func (q *CatDogQueue) PollDog() Dog {
    if q.IsDogEmpty() {
        panic("No dogs")
    }

    dog := q.dogs[0].dog
    q.dogs = q.dogs[1:]

    return dog
}

func (q *CatDogQueue) PollAll() Pet {
    if q.IsEmpty() {
        panic("No pets")
    }

    if q.IsDogEmpty() {
        return q.PollCat()
    } else if q.IsCatEmpty() {
        return q.PollDog()
    }

    dog := q.dogs[0]
    cat := q.cats[0]

    if dog.seq < cat.seq {
        return q.PollDog()
    } else {
        return q.PollCat()
    }
}

func NewCatDogQueue() *CatDogQueue {
    return &CatDogQueue{ seq : 0 }
}


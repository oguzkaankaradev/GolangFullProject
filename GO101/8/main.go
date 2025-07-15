//MixedCaps : isimlendime kuralları
//Golangde isimlendirme standarlrı çok kelimeli  kullanı aşağıdaki gibidir.
//userName(ilk harf küçük ikinci ve sonraki kelimelerin harfleri büyük ve bitişik.)
//UserName (Her kelimemenin başlangıcı büyük ve bitişiktir.)

//Yanlış kullanım
//user_name gget_usernamerGetUserId

//Bu kural neden önnemlidir.
//Kodunuz daha okunabilir hale gelir, sizden sonraki go lang yazılımcısı veya ekibe katılan yeni bir developer kolaylıkla entegre olur
//Bu kurallara MixedCaps denilmektedir.

type BlogPost struct {
	AuthorName string
}

type Blog_Post struct {
	AuthorName string
}
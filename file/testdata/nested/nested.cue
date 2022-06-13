package nested

#Continent: "Europe" | "Asia" | "Africa" | "North America" | "South America" | "Australia" | "Antarctica"

#Tree: {
	size: >0

	age: >0

	location: #Continent | [...#Continent]
}

Oak: #Tree & {
	size: 30

	age: 700

	location: "Europe" | "North America" | "Australia"
}

Bonsai: #Tree & {
	size: 3

	age: 1000

	location: "Asia"
}

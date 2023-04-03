package main;


func listen(){
	// Escuchamos quejas de esclavos
	// Tiene que pasar un minimo
	// TODO: Tener una variable y ver cuanto seria el minimo, 60%?
	// TODO: Leer el procentaje desde entonro
}

func heartbeat(){
	// Revisar cuantos nodos tenemos activos para el minimo
}

func communicate(){
	// Llamado despues de decide
	// Nombramos quien es el nuevo master y quienes siguen esclavos
}

func decide(){
	// Elegir uno de los nodos bajo algun criterio
	// LLama heartbeat para saber con quien decidir
	// Elegimos random - El queue mas largo!!!!
}
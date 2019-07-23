// Package mongoapi sse é apenas um exemplo de API simples para conexão com o banco de dados
// MongoDB e que contém operações CRUD
package mongoapi

import (
	"context"
	"fmt"
	"log"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

// User é um tipo de define um usuário do sistema
type User struct {
	Name  string `json:"name,omitempty"`
	Email string `json:"email,omitempty"`
	Age   int    `json:"age,omitempty"`
}

/*
func main() {

	// Data base conection
	URI := "mongodb://localhost:27017"
	client := IniciarConexao(URI)

	// Close the conection at the end of the program
	defer EncerrarConexao(*client)

	// Cria o caminho da coleção de interesse
	collection := client.Database("golang-test").Collection("users")

	// Instancia um novo usuário
	usuario := User{"Usuario 1", "usuarioum@mail.com", 18}
	fmt.Println(usuario)

	//Inserção
	InserirUsuario(*client, *collection, usuario)

	// Consulta
	filtro := bson.D{{"name", "Usuario 1"}}

	usuarioConsultado := ConsultarUsuario(*client, *collection, filtro)
	fmt.Println(usuarioConsultado)

	//filtro = bson.D{{"age", bson.D{{"$gte", 18}}}}
	filtro = bson.D{{}} //todos os elementos

	listaUsuarios := ConsultarUsuarios(*client, *collection, filtro)

	for _, usuario := range listaUsuarios {
		fmt.Println(usuario)
	}

	// Atualizaçao
	filtro = bson.D{{"age", bson.D{{"$gte", 30}}}}
	update := bson.D{{"$set", bson.D{{
		"name", "DELL"}}}}

	AtualizarUsuario(*client, *collection, filtro, update)

	//Delete
	filtro = bson.D{{"name", "Ricardo"}}
	ExcluirUsuario(*client, *collection, filtro)

	ZerarColecao(*client, *collection)

}
*/

// IniciarConexao inicia uma conexão com o banco de dados.
func IniciarConexao(uri string) *mongo.Client {
	clientOptions := options.Client().ApplyURI(uri)
	client, err := mongo.Connect(context.TODO(), clientOptions)

	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Conexão com o MongoDB estabelecida com sucesso!")

	return client
}

// EncerrarConexao encerra a conexao com o banco de dados.
func EncerrarConexao(client mongo.Client) {
	err := client.Disconnect(context.TODO())

	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Conexão com o MongoDB encerrada com sucesso!")
}

// InserirUsuario insere um usuário no banco de dados.
func InserirUsuario(client mongo.Client, collection mongo.Collection, usuario User) {
	result, err := collection.InsertOne(context.TODO(), usuario)

	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Usuário inserido com sucesso! ", result.InsertedID)
}

// ConsultarUsuario retorna um usuário de acordo com a pesquisa.
func ConsultarUsuario(client mongo.Client, collection mongo.Collection, consulta bson.D) User {
	var result User
	err := collection.FindOne(context.TODO(), consulta).Decode(&result)

	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Usuário consultado com sucesso! ", result)
	return result
}

// ConsultarUsuarios retorna um slice de usuarios de acordo com a consulta.
func ConsultarUsuarios(client mongo.Client, collection mongo.Collection, consulta bson.D) []User {
	cursor, err := collection.Find(context.TODO(), consulta)

	if err != nil {
		log.Fatal(err)
	}

	var usuarios []User

	for cursor.Next(context.TODO()) {
		var usuario User

		// decode the document
		err := cursor.Decode(&usuario)

		if err != nil {
			log.Fatal(err)
		}

		usuarios = append(usuarios, usuario)
	}

	fmt.Println("Usuários consultados com sucesso!")
	return usuarios
}

// AtualizarUsuario atualiza o registro do primeiro usuario encontrado na consulta,
// de acordo com o a atualizaçao.
func AtualizarUsuario(client mongo.Client, collection mongo.Collection, consulta bson.D, atualizacao bson.D) {

	result, err := collection.UpdateMany(context.TODO(), consulta, atualizacao)

	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Usuário atualizado com sucesso! ", result)
}

// ExcluirUsuario apaga o registro de um usuário no banco de dados.
func ExcluirUsuario(client mongo.Client, collection mongo.Collection, consulta bson.D) {

	result, err := collection.DeleteOne(context.TODO(), consulta)

	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Usuário deletado com sucesso! ", result)
}

// ZerarColecao apaga todos os registros presentes na celecao.
func ZerarColecao(client mongo.Client, collection mongo.Collection) {
	consulta := bson.D{{}}
	result, err := collection.DeleteMany(context.TODO(), consulta)

	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Todos os registros do banco de dados foram apagados", result)
}

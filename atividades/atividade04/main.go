package main

import (
	"fmt"
)

type Veiculo struct {
	Placa          string
	Modelo         string
	Ano            int
	Abastecimentos []float64
}
type Frota struct {
	TodosVeiculos []Veiculo
}

func (f *Frota) AdicionarVeiculo(veiculo Veiculo) {
	f.TodosVeiculos = append(f.TodosVeiculos, veiculo)
}

func (f Frota) ListarVeiculos() []Veiculo {
	return f.TodosVeiculos
}

func (f Frota) ValidarAno(AnoRecebido int) bool {
	if AnoRecebido >= 1900 && AnoRecebido <= 2025 {
		return true
	}
	return false
}

func (f Frota) PlacaExiste(placaRecebida string) bool {
	for _, veiculo := range f.TodosVeiculos {
		if veiculo.Placa == placaRecebida {
			return true
		}
	}
	return false
}

// Tenta registrar um novo abastecimento ao veículo pela placa.
// Valida placa e valor. Se tudo certo, adiciona ao slice de abastecimentos.
func (f Frota) RegistrarAbastecimento(placaRecebida string, valor float64) {
	if !f.PlacaExiste(placaRecebida) {
		fmt.Printf("Placa não cadastrada: %s", placaRecebida)
		return
	}
	if valor <= 0 {
		fmt.Println("O valor do abastecimento deve ser maior que 0")
	}
	return

	// Percorre todos os veículos da frota para encontrar o correto.
	for i, v := range f.TodosVeiculos {
		if v.Placa == placaRecebida {
			// Adiciona o valor ao slice de abastecimentos do veículo encontrado.
			f.TodosVeiculos[i].Abastecimentos = append(f.TodosVeiculos[i].Abastecimentos, valor)
		}
	}

	fmt.Printf("O abastecimento de %.2f no veículo %s foi feito com sucesso!", placaRecebida, valor)
}

func (f Frota) ResumoAbastecimento() []Veiculo {
	return f.TodosVeiculos
}

func main() {

	var MinhaFrota Frota // Cria uma única variável da frota que será usada durante toda a execução do programa

	for {

		fmt.Println("----MENU FROTA----")
		fmt.Println("1 - Adicionar Veiculo")
		fmt.Println("2 - Registrar Abastecimento")
		fmt.Println("3 - Listar Veiculos")
		fmt.Println("4 - Mostrar Resumo de Abastecimentos")
		fmt.Println("5 - Sair \n")

		fmt.Print("ESCOLHA UMA OPÇÃO: ")

		var opcao string
		fmt.Scan(&opcao)

		switch opcao {

		case "1": // [Cadastro] Solicita dados, valida placa/ano, adiciona veículo novo à frota

			var Modelo string
			var Placa string
			var Ano int

			fmt.Print("Qual o modelo do Carro: ")
			fmt.Scanln(&Modelo)

			for {
				fmt.Print("Qual a placa do Carro: ")
				fmt.Scanln(&Placa)
				valido := MinhaFrota.PlacaExiste(Placa)

				if !valido {
					break
				} else {
					fmt.Println("A placa é duplicada, tente outra!")
				}
			}

			for {
				fmt.Print("Qual o ano do Carro: ")
				fmt.Scanln(&Ano)

				Valido := MinhaFrota.ValidarAno(Ano)

				if Valido {
					break
				} else {
					fmt.Println("O ano informado não é valido, tente novamente!")
				}
			}
			novoVeiculo := Veiculo{
				Modelo:         Modelo,
				Placa:          Placa,
				Ano:            Ano,
				Abastecimentos: []float64{},
			}

			MinhaFrota.AdicionarVeiculo(novoVeiculo)
			fmt.Println("O veiculo foi adicionado com sucesso! \n")

		case "2": // [Abastecimento] Solicita placa existente e valor, adiciona abastecimento ao veículo correto.
			var placa string

			for {
				fmt.Print("Qual é a placa do carro: ")
				fmt.Scanln(&placa)

				if MinhaFrota.PlacaExiste(placa) {
					break
				} else {
					fmt.Println("Placa não encontrada, tente outra placa!")
				}
			}

			var valor float64

			for {
				fmt.Print("Qual o valor de combustível que você quer colocar: ")
				fmt.Scanln(&valor)

				if valor >= 0.0 {
					fmt.Println("Gasolina colocada com sucesso \n")
					break
				} else {
					fmt.Println("O valor não é válido, tente novamente!")
				}
			}

			for i, veiculo := range MinhaFrota.TodosVeiculos {
				if veiculo.Placa == placa {
					// Aqui encontrado: use o índice i
					MinhaFrota.TodosVeiculos[i].Abastecimentos = append(
						MinhaFrota.TodosVeiculos[i].Abastecimentos, valor,
					)
					break
				}
			}

		case "3": // [Listar] Percorre e mostra todos os veículos cadastrados com seus abastecimentos.

			for _, veiculo := range MinhaFrota.ListarVeiculos() {
				fmt.Printf("Modelo: %s, Placa: %s, Ano: %d, Abastecimento: %.2f \n", veiculo.Modelo, veiculo.Placa, veiculo.Ano, veiculo.Abastecimentos)
			}

		case "4": // [Resumo] Solicita placa, exibe apenas os abastecimentos do veículo escolhido.

			var placa string

			for {
				fmt.Print("Você quer consultar o resumo de qual veiculo: ")
				fmt.Scanln(&placa)

				if MinhaFrota.PlacaExiste(placa) {
					break
				} else {
					fmt.Println("Placa não encontrada, tente outra placa!")
				}
			}

			for _, veiculo := range MinhaFrota.ResumoAbastecimento() {
				if veiculo.Placa == placa {
					fmt.Printf("Placa: %s, Abastecimento: %.2f \n", veiculo.Placa, veiculo.Abastecimentos)
				}
			}
		case "5":
		}
	}
}

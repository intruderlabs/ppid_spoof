package main

import (
	"fmt"
	"os"
	"os/exec"
	"syscall"
)

const (
	PROCESS_CREATE_PROCESS = 0x0080
	MAXIMUM_ALLOWED        = 0x02000000
)

func getProcessHandle(pid uint32) (syscall.Handle, error) {
	handle, err := syscall.OpenProcess(MAXIMUM_ALLOWED, false, pid)
	if err != nil {
		return 0, fmt.Errorf("falha ao abrir o processo pai (PID %d): %w", pid, err)
	}
	return handle, nil
}

func main() {
	var parentPid uint32 = 16824 // <-- SUBSTITUA PELO PID REAL DO PROCESSO PAI DESEJADO

	parentHandle, err := getProcessHandle(parentPid)
	if err != nil {
		fmt.Printf("Erro: %v\n", err)
		os.Exit(1)
	}
	defer syscall.CloseHandle(parentHandle)

	cmd := exec.Command("C:\\Windows\\System32\\notepad.exe")

	cmd.SysProcAttr = &syscall.SysProcAttr{
		ParentProcess: parentHandle, // <-- Aqui definimos o "pai" falso
	}

	fmt.Printf("[*] Iniciando notepad.exe com PPID spoofado para o PID %d...\n", parentPid)
	if err := cmd.Start(); err != nil {
		fmt.Printf("Erro ao iniciar o processo: %v\n", err)
		os.Exit(1)
	}

	fmt.Printf("[+] Processo iniciado com sucesso! PID real: %d\n", cmd.Process.Pid)
	fmt.Println("[*] Verifique o Gerenciador de Tarefas ou o Process Explorer para confirmar.")
	fmt.Println("    O processo notepad.exe deve aparecer como filho do processo especificado.")

}

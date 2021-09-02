is_run = True
n_loker = 0

loker = []

def empty_n_loker():
	if n_loker == 0:
		print("n loker masih kosong. (harus init value)")
		return True
	
	return False

while is_run:
	argumen = ""
	command = input("> ")

	if " " in command:
		split_command = command.split()

		if len(split_command) > 2:
			command = split_command[0]
			argumen = [split_command[1], split_command[2]]
		else:
			command, argumen = split_command

		# print(argumen)

	if command == "init":
		try:
			n_loker = int(argumen)
			if n_loker > 0:
				print("Berhasil membuat loker dengan jumlah " + str(n_loker))

				for i in range(n_loker):
					temp = {"nomor": str(i+1), "tipe_id": "-", "nomor_id": "-", "status": "kosong"}
					loker.append(temp)
			else:
				print("n loker tidak boleh 0 atau kurang dari 0")
		except:
			print("Invalid Value")

	elif command == "exit":
		is_run = False

	elif command == "input":
		if argumen == "":
			print("Invalid Argumen")

		elif not empty_n_loker():
			try:
				kosong = False
				for l in loker:
					if l['status'] == "kosong":
						l['tipe_id'] = argumen[0]
						l['nomor_id'] = argumen[1]
						l['status'] = "ada"
						kosong = True

						print("Kartu identitas tersimpan di loker nomor " + l['nomor'])
						break

				if not kosong:
					print("Maaf loker sudah penuh")

			except:
				print("Invalid Argumen")

	elif command == "status":
		if not argumen == "":
			print("Invalid Argumen")

		elif not empty_n_loker():
			print("No Loker\tTipe Identitas\tNo Identitas")
			for l in loker:
				print(l['nomor'], end="\t\t")
				print(l['tipe_id'], end="\t\t")
				print(l['nomor_id'])

	elif command == "leave":
		if argumen == "":
			print("Invalid Argumen")
			
		elif not empty_n_loker():
			try:
				ketemu = False

				for l in loker:
					if l['nomor'] == argumen:
						l['status'] = "kosong"
						l['tipe_id'] = "-"
						l['nomor_id'] = "-"

						print("Loker nomor " + str(l['nomor']) + " berhasil dikosongkan")
						ketemu = True

				if not ketemu:
					print("Loker tidak ditemukan!")

			except:
				print("Invalid Argumen")

	elif command == "find":
		if argumen == "":
			print("Invalid Argumen")

		elif not empty_n_loker():
			try:
				ketemu = False

				for l in loker:
					if l['nomor_id'] == argumen:
						print("Kartu identitas tersebut berada di loker nomor " + l['nomor'])							
						ketemu = True

				if not ketemu:
					print("Nomor identitas tidak ditemukan")

			except:
				print("Invalid Argumen")

	elif command == "search":
		if argumen == "":
			print("Invalid Argumen")

		elif not empty_n_loker():
			try:
				ketemu = False
				result = ""

				for l in loker:
					if l['tipe_id'] == argumen:
						result += l['nomor_id'] + ','
						ketemu = True

				if not ketemu:
					print("Loker tidak ditemukan!")
				else:
					result.rstrip(',')
					print(result)

			except:
				print("Invalid Argumen")

	else:
		print("Invalid Command")
		print("Perintah tidak ditemukan")
		print()
	# print(loker)

print("Program berakhir")


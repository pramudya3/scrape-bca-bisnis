package scraper

import (
	"log"
	"strconv"
	"strings"
	"time"

	"github.com/playwright-community/playwright-go"
	"go.mau.fi/libsignal/logger"
)

const baseBCABisnis = "https://vpn.klikbca.com/+CSCOE+/logon.html"

func loginBCABisnis(page playwright.Page, corpID, username, password string) playwright.Page {

	// login ke-1
	if _, err := page.Goto(baseBCABisnis); err != nil {
		log.Println("Gagal menuju website BCA Bisnis :", err)
	}

	userName := corpID + username
	userName = strings.TrimSpace(userName)

	usernameField, err := page.WaitForSelector("input[id='username']")
	if err != nil {
		log.Println("Gagal mendapatkan element usernameField :", err)
	} else {
		if err = usernameField.Type(userName); err != nil {
			log.Println("Gagal mengetik username :", err)
		} else {
			log.Println("berhasil mengetik username :", userName)
		}
	}

	passwordField, err := page.Locator("input[type='password']")
	if err != nil {
		log.Println("Gagal mendapatkan element passwordField :", err)
	}
	if err = passwordField.Type(password); err != nil {
		log.Println("Gagal mengetik password :", err)
	} else {
		log.Println("mengetik password")
	}

	btnLogin, err := page.WaitForSelector("input#submit.btn.text-white.vpn-btn.mt-2.mb-3")
	if err != nil {
		log.Println("Gagal mendapatkan element btnLogin :", err)
	} else {
		page.WaitForTimeout(1000)
		if err = btnLogin.Click(); err != nil {
			log.Println("Gagal klik btnLogin :", err)
		} else {
			log.Println("Klik btnLogin 1")
		}
	}

	page.WaitForTimeout(2000)
	KlikBCABisnis, err := page.WaitForSelector("body > div.main > div > div > div.col-12.col-sm-12.col-md-12.col-lg-8.col-xl-6 > div.row.mt-6vh.mb-3 > div:nth-child(1) > a")
	if err != nil {
		log.Println("Gagal mendapatkan element KlikBCABisnis :", err)
	} else {
		page.WaitForTimeout(1000)
		if err = KlikBCABisnis.Click(); err != nil {
			log.Println("Gagal klik KlikBCABisnis :", err)
		}
	}

	return page
}

func loginBCABisnis2(page playwright.Page, corpId, username, password string) playwright.Page {
	// login ke-2
	corpIdfield, err := page.WaitForSelector("input[name='corp_cd']")
	if err != nil {
		log.Println("Gagal mendapatkan element corpIdField :", err)
	} else {
		if err = corpIdfield.Type(corpId); err != nil {
			log.Println("Gagal mengetik corpID :", err)
		} else {
			log.Println("mengetik corpID")
		}
	}

	userNameField, err := page.WaitForSelector("input[name='user_cd']")
	if err != nil {
		log.Println("Gagal mendapatkan element userNameField :", err)
	} else {
		if err = userNameField.Type(username); err != nil {
			log.Println("Gagal mengetik username :", err)
		} else {
			log.Println("mengetik username")
		}
	}

	passWordField, err := page.WaitForSelector("input[name='pswd']")
	if err != nil {
		log.Println("Gagal mendapatkan element passWordField :", err)
	} else {
		if err = passWordField.Type(password); err != nil {
			log.Println("Gagal mengetik password :", err)
		} else {
			log.Println("mengetik password ")
		}
	}

	btnLogin2, err := page.WaitForSelector("img[name='Image13']")
	if err != nil {
		log.Println("Gagal mendapatkan element btnLogin2 :", err)
	} else {
		if err = btnLogin2.Click(); err != nil {
			log.Println("Gagal klik btnLogin2 :", err)
		} else {
			log.Println("klik btnLogin 2, Login BCA Bisnis")
		}
	}

	return page
}

func menuBCABisnis(page playwright.Page, nomorRekening string, totalCekHari int64) playwright.Page {

	leftFrame, err := page.WaitForSelector("frame[name='leftFrame']")
	if err != nil {
		log.Println("Gagal mendapatkan element leftFrame :", err)
	}

	leftFrameContent, err := leftFrame.ContentFrame()
	if err != nil {
		log.Println("Gagal mendapat element leftFrameContent :", err)
	}

	informasiRekening, err := leftFrameContent.WaitForSelector("div#divFold1 a.clFoldLinks")
	if err != nil {
		log.Println("Gagal mendapatkan element informasiRekening :", err)
	} else {
		if err = informasiRekening.Click(); err != nil {
			log.Println("Gagal klik inforamsiRekening :", err)
		} else {
			log.Println("Klik informasi rekening")
		}
	}

	informasiSaldo, err := leftFrameContent.WaitForSelector("div#divFoldSub1_0.clFoldSub a.clSublinks")
	if err != nil {
		log.Println("Gagal mendapat element informasiSaldo :", err)
	} else {
		if err = informasiSaldo.Click(); err != nil {
			log.Println("Gagal klik informasiSaldo :", err)
			page.WaitForTimeout(1000)
		}
	}

	workspaceFrame, err := page.WaitForSelector("frame[name='workspace']")
	if err != nil {
		log.Println("Gagal mendpatkan workspaceFrame :", err)
	}

	workspaceFrameContent, err := workspaceFrame.ContentFrame()
	if err != nil {
		log.Println("Gagal mendapatkan element workspaceFrameContent :", err)
	}
	page.WaitForTimeout(2000)

	checkAll, err := workspaceFrameContent.WaitForSelector("input[name='allAcct']")
	if err != nil {
		log.Println("Gagal mendapatkan element checkAll :", err)
	} else {
		if err = checkAll.Check(); err != nil {
			log.Println("Gagal checkAll :", err)
		} else {
			log.Println("checked all rekening")
		}
	}

	btnKirim, err := workspaceFrameContent.WaitForSelector("input[name='Submit']")
	if err != nil {
		log.Println("Gagal mendapatkan element btnKirim :", err)
	} else {
		if err = btnKirim.Click(); err != nil {
			log.Println("Gagal klik btnKirim :", err)
			page.WaitForTimeout(2000)
		} else {
			log.Println("btnKirim diklik, mulai scraping saldo rekening")
		}
	}

	workspaceFrame, err = page.WaitForSelector("frame[name='workspace']")
	if err != nil {
		log.Println("Gagal mendpatkan workspaceFrame :", err)
	}

	workspaceFrameContent, err = workspaceFrame.ContentFrame()
	if err != nil {
		log.Println("Gagal mendapatkan element workspaceFrameContent :", err)
	}
	page.WaitForTimeout(2000)

	tabelRekenings, err := workspaceFrameContent.QuerySelectorAll("table.clsform") //tr.clsEven
	if err != nil {
		log.Println("Gagal mendapat element tabelRekenings :", err)
	}
	page.WaitForTimeout(1000)

	for k := 0; k < len(tabelRekenings); k++ {

		// table rek 1
		tableRek1, err := workspaceFrameContent.QuerySelectorAll("#frmParam > table:nth-child(52) tr.clsEven")
		if err != nil {
			log.Println("Gagal mendapat element tableRek1 :", err)
		}

	myRekening1:
		for _, datas1 := range tableRek1 {
			cells, err := datas1.QuerySelectorAll("td")
			if err != nil {
				log.Println("gagal mendapatkan element cells :", err)
			}

			for q, cell := range cells {
				cellValue, err := cell.InnerText()
				if err != nil {
					log.Println("Gagal mendapatkan element cellValue :", err)
				}

				switch q {
				case 0:
					log.Println("case 0 :", cellValue)
					caseZero := strings.Replace(cellValue, "-", "", -1)
					log.Println("caseZero :", caseZero)
					log.Println("nomor rekening :", nomorRekening)
					if caseZero != nomorRekening {
						break myRekening1
					}
				case 1:
					log.Println("case 1 :", cellValue)
				case 2:
					log.Println("case 2 :", cellValue)
					saldo := strings.Replace(cellValue, "Rp", "", -1)
					saldo = strings.Replace(saldo, ",", "", -1)
					saldo = strings.ReplaceAll(saldo, "\u00a0", "")
					saldo = strings.ReplaceAll(saldo, "&nbsp;", "")
					caseTwo := strings.Split(saldo, ".")
					saldoInt, err := strconv.Atoi(caseTwo[0])
					if err != nil {
						log.Println("Gagal parse string to int :", err)
					}
					log.Println(saldoInt)
				}
			}
		}

		// table rek 2

		tableRek2, err := workspaceFrameContent.QuerySelectorAll("#frmParam > table:nth-child(54) tr.clsEven")
		if err != nil {
			log.Println("Gagal mendapatkan element tableRek2 :", err)
		}

	myRekening2:
		for _, datas := range tableRek2 {

			cells, err := datas.QuerySelectorAll("td")
			if err != nil {
				log.Println("Gagal mendapatkan element cell :", err)
			}

			for a, cell := range cells {
				cellValue, err := cell.InnerText()
				if err != nil {
					log.Println("Gagal mendapatkan element cellValue :", err)
				}

				switch a {
				case 1:
					log.Println("case 1 :", cellValue)
					caseOne := strings.Replace(cellValue, "-", "", -1)
					if caseOne != nomorRekening {
						break myRekening2
					}
				case 2:
					log.Println("case 2 :", cellValue)
				case 3:
					log.Println("case 3 :", cellValue)
					saldo := strings.Replace(cellValue, "Rp", "", -1)
					saldo = strings.Replace(saldo, ",", "", -1)
					saldo = strings.ReplaceAll(saldo, "\u00a0", "")
					saldo = strings.ReplaceAll(saldo, "&nbsp;", "")
					caseThree := strings.Split(saldo, ".")
					saldoInt, err := strconv.Atoi(caseThree[0])
					if err != nil {
						log.Println("Gagal parse string to int :", err)
					}
					log.Println(saldoInt)
				}
			}
		}
	}

	//menu mutasi
	informasiMutasi, err := leftFrameContent.WaitForSelector("div#divFoldSub1_1.clFoldSub a.clSubLinks")
	if err != nil {
		log.Println("Gagal mendapat element informasiMutasi :", err)
	} else {
		page.WaitForTimeout(1000)
		if err = informasiMutasi.Click(); err != nil {
			log.Println("Gagal klik informasiMutasi :", err)
		} else {
			log.Println("klik mutasi rekening")
		}
	}

	workspaceFrame, err = page.WaitForSelector("frame[name='workspace']")
	if err != nil {
		log.Println("Gagal mendpatkan workspaceFrame :", err)
	}

	workspaceFrameContent, err = workspaceFrame.ContentFrame()
	if err != nil {
		log.Println("Gagal mendapatkan element workspaceFrameContent :", err)
	}
	page.WaitForTimeout(2000)
	log.Println("sukses")

	days := time.Duration(totalCekHari)
	start := time.Now().Add(-days * 24 * time.Hour)
	startDay := start.Format("02")

	fromDay, err := workspaceFrameContent.WaitForSelector("select#from_day")
	if err != nil {
		log.Println("Gagal mendapat element fromDay :", err)
	} else {
		if _, err = fromDay.SelectOption(playwright.SelectOptionValues{Values: playwright.StringSlice(startDay)}); err != nil {
			log.Println("Gagal select date :", err)
		} else {
			log.Println("berhasil memilih tanggal mulai scraping mutasi")
		}
	}

	searchRekening, err := workspaceFrameContent.WaitForSelector("input#acct_display")
	if err != nil {
		log.Println("Gagal mendapatkan element searchRekening :", err)
	} else {
		if err = searchRekening.Type(nomorRekening); err != nil {
			log.Println("Gagal klik searchRekening :", err)
		}

		page.WaitForTimeout(1000)
		if err = searchRekening.Press(`Enter`); err != nil {
			log.Println("Gagal press enter :", err)
		}
	}

	_, err = workspaceFrameContent.WaitForSelector("h3.clsErrorMsg", playwright.PageWaitForSelectorOptions{Timeout: playwright.Float(5000)})
	if err != nil {
		log.Println("MUTASI DITEMUKAN")

		tahun, err := workspaceFrameContent.WaitForSelector("body > form > table:nth-child(1) > tbody > tr:nth-child(3) > td:nth-child(3)")
		if err != nil {
			log.Println("Gagal mendapatkan element tahun :", err)
		}
		tahunString, err := tahun.InnerText()
		if err != nil {
			log.Println("Gagal mendapatkan element tahunString :", err)
		}

		year := strings.Split(tahunString, "/")
		year[4] = strings.TrimSpace(year[4])

		workspaceFrame, err = page.WaitForSelector("frame[name='workspace']")
		if err != nil {
			log.Println("Gagal mendpatkan workspaceFrame :", err)
		}

		page.WaitForTimeout(2000)
		workspaceFrameContent, err = workspaceFrame.ContentFrame()
		if err != nil {
			log.Println("Gagal mendapatkan element workspaceFrameContent :", err)
		}

		nama, err := workspaceFrameContent.WaitForSelector("body > form > table:nth-child(1) > tbody > tr:nth-child(2) > td:nth-child(3)")
		if err != nil {
			log.Println("Gagal mendapatkan element nama :", err)
		}
		namaString, err := nama.InnerText()
		if err != nil {
			log.Println("Gagal parse to string :", err)
		}
		pemilik := strings.Replace(namaString, ":", "", -1)
		pemilik = strings.ReplaceAll(pemilik, "&nbsp;", "")

		btnNext, err := workspaceFrameContent.Locator("input#Next")
		if err != nil {
			log.Println("Gagal mendapatkan element btnNext :", err)
		}
		page.WaitForTimeout(1000)

		isVisible, err := btnNext.IsVisible()
		if err != nil {
			log.Println("Gagal mendapatkan state element btnNext :", err)
		}

		var tableMutasi []playwright.ElementHandle

		for play := true; play; play = isVisible {

			workspaceFrame, err := page.WaitForSelector("frame[name='workspace']")
			if err != nil {
				log.Println("Gagal mendpatkan workspaceFrame :", err)
			}
			page.WaitForTimeout(2000)

			workspaceFrameContent, err := workspaceFrame.ContentFrame()
			if err != nil {
				log.Println("Gagal mendapatkan element workspaceFrameContent :", err)

			}

			tableMutasi, err = workspaceFrameContent.QuerySelectorAll("table.clsForm tr[class]")
			if err != nil {
				log.Println("Gagal mendapatkan element tableMutasi :", err)
			}

		myLoop:
			for _, rowMutasi := range tableMutasi {

				cells, err := rowMutasi.QuerySelectorAll("td")
				if err != nil {
					log.Println("Gagal mendapatkan element cell :", err)
				}

				for a, cell := range cells {
					mutasiCell, err := cell.InnerText()
					if err != nil {
						log.Println("Gagal mendapat element mutasiCell :", err)
					}

					switch a {
					case 0:
						if mutasiCell != "PEND" {
							now := time.Now()
							log.Println(now)

						} else {
							tanggalBank := mutasiCell + year[4]
							log.Println("tanggalBank :", tanggalBank)
						}
					case 1:
						ket := strings.Replace(mutasiCell, "\n", "", -1)
						log.Println("case 1 :", ket)
						if mutasiCell == "" || ket == "Saldo Awal" || ket == "" {
							continue myLoop
						} else {
							log.Println(mutasiCell)
						}
					case 3:
						jumlah := strings.Split(mutasiCell, ".")
						jumlah[0] = strings.Replace(jumlah[0], ",", "", -1)
						jumlah[0] = strings.ReplaceAll(jumlah[0], "&nbsp;", "")
						jumlahInt, err := strconv.Atoi(jumlah[0])
						jumlahType := strings.Split(jumlah[1], " ")
						if jumlahType[1] == "CR" {
							log.Println("KREDIT")
						} else {
							log.Println("DEBET")
						}
						if err != nil {
							log.Println("Gagal parsing string ke int :", err)
						}
						log.Println("case 3 :", jumlahInt)
					case 4:
						saldo := strings.Split(mutasiCell, ".")
						saldo[0] = strings.Replace(saldo[0], ",", "", -1)
						saldo[0] = strings.ReplaceAll(saldo[0], "&nbsp;", "")
						saldo[0] = strings.ReplaceAll(saldo[0], "\u00a0", "")
						jumlahInt, err := strconv.Atoi(saldo[0])
						if err != nil {
							log.Println("Gagal parsing string ke int :", err)
						}
						log.Println("case 4 :", jumlahInt)
					}
				}
			}

			btnNext, err = workspaceFrameContent.Locator("input#Next")
			if err != nil {
				break
			} else {
				if isVisible {
					if err = btnNext.Click(); err != nil {
						log.Println("Gagal klik btnNext :", err)
					} else {
						log.Println("Halaman selanjutnya")
					}
				}
			}
		}

	} else {
		log.Println("MUTASI TIDAK DITEMUKAN")
	}

	closeInformasiRekening, err := leftFrameContent.WaitForSelector("div#divFold1 a.clFoldLinks")
	if err != nil {
		log.Println("Gagal mendapat element closeInformasiRekening :", err)
	} else {
		if err = closeInformasiRekening.Click(); err != nil {
			log.Println("Gagal klik closeInformasiRekening :", err)
		}
	}

	backToHomeBtn, err := leftFrameContent.WaitForSelector("div#divFold0 a.clFoldLinks")
	if err != nil {
		log.Println("Gagal mendapatkan element backToHomeBtn :", err)
	} else {
		if err = backToHomeBtn.Click(); err != nil {
			log.Println("Gagal klik backToHomeBtn :", err)
		} else {
			log.Println("Back To Home")
		}
	}

	return page
}

func isBCABisnisLogin(page playwright.Page) bool {
	if _, err := page.WaitForSelector("input#username", playwright.PageWaitForSelectorOptions{Timeout: playwright.Float(5000)}); err != nil {
		return true
	} else {
		return false
	}
	return true
}

func logoutBCABisnis(page playwright.Page) {
	topFrame, err := page.WaitForSelector("frame[name='topFrame']")
	if err != nil {
		log.Println("Gagal mendapatkan element topFrame :", err)
	}

	topFrameContent, err := topFrame.ContentFrame()
	if err != nil {
		log.Println("Gagal mendapat element topFrameContent :", err)
	}

	btnLogout, err := topFrameContent.WaitForSelector("img[name='logout']")
	if err != nil {
		log.Println("Gagal mendapat element btnLogout :", err)
	} else {
		if err = btnLogout.Click(); err != nil {
			log.Println("Gagal klik btnLogout :", err)
		}
	}
}

func main() {
	var (
		corpid, username, password, norek string
		totalCekHari                      int64
	)

	pw, err := playwright.Run()
	if err != nil {
		log.Println(err)
	}

	browser, err := pw.WebKit.Launch()
	if err != nil {
		log.Println(err)
	}

	page, err := browser.NewPage()
	if err != nil {
		logger.Debug(err)
	}

	page = loginBCABisnis(page, corpid, username, password)
	page = loginBCABisnis2(page, corpid, username, password)
	page = menuBCABisnis(page, norek, totalCekHari)
	logoutBCABisnis(page)
}

package push

import (
	"io/ioutil"
	"os"
	"path/filepath"

	"code.cloudfoundry.org/cli/integration/helpers"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	. "github.com/onsi/gomega/gexec"
)

var _ = Describe("push with symlink path", func() {
	var (
		appName       string
		runningDir    string
		symlinkedPath string
	)

	BeforeEach(func() {
		appName = helpers.NewAppName()

		var err error
		runningDir, err = ioutil.TempDir("", "push-with-symlink")
		Expect(err).ToNot(HaveOccurred())
		symlinkedPath = filepath.Join(runningDir, "symlink-dir")
	})

	AfterEach(func() {
		Expect(os.RemoveAll(runningDir)).To(Succeed())
	})

	Context("push with flag options", func() {
		When("pushing from a symlinked current directory", func() {
			It("should push with the absolute path of the app", func() {
				helpers.WithHelloWorldApp(func(dir string) {
					Expect(os.Symlink(dir, symlinkedPath)).ToNot(HaveOccurred())

					session := helpers.CustomCF(helpers.CFEnv{WorkingDirectory: symlinkedPath}, PushCommandName, appName)
					Eventually(session).Should(helpers.SayPath(`path:\s+%s`, dir))
					Eventually(session).Should(Exit(0))
				})
			})
		})

		When("pushing a symlinked path with the '-p' flag", func() {
			It("should push with the absolute path of the app", func() {
				helpers.WithHelloWorldApp(func(dir string) {
					Expect(os.Symlink(dir, symlinkedPath)).ToNot(HaveOccurred())

					session := helpers.CF(PushCommandName, appName, "-p", symlinkedPath)
					Eventually(session).Should(helpers.SayPath(`path:\s+%s`, dir))
					Eventually(session).Should(Exit(0))
				})
			})
		})

		When("pushing an symlinked archive with the '-p' flag", func() {
			var archive string

			BeforeEach(func() {
				helpers.WithHelloWorldApp(func(appDir string) {
					tmpfile := helpers.TempFileAbsolutePath("", "push-archive-integration")
					archive = tmpfile.Name()
					Expect(tmpfile.Close()).ToNot(HaveOccurred())

					err := helpers.Zipit(appDir, archive, "")
					Expect(err).ToNot(HaveOccurred())
				})
			})

			AfterEach(func() {
				Expect(os.RemoveAll(archive)).ToNot(HaveOccurred())
			})

			It("should push with the absolute path of the archive", func() {
				Expect(os.Symlink(archive, symlinkedPath)).ToNot(HaveOccurred())

				session := helpers.CF(PushCommandName, appName, "-p", symlinkedPath)
				Eventually(session).Should(helpers.SayPath(`path:\s+%s`, archive))
				Eventually(session).Should(Exit(0))
			})
		})

		Context("push with a single app manifest", func() {
			When("the path property is a symlinked path", func() {
				It("should push with the absolute path of the app", func() {
					helpers.WithHelloWorldApp(func(dir string) {
						Expect(os.Symlink(dir, symlinkedPath)).ToNot(HaveOccurred())

						helpers.WriteManifest(filepath.Join(runningDir, "manifest.yml"), map[string]interface{}{
							"applications": []map[string]string{
								{
									"name": appName,
									"path": symlinkedPath,
								},
							},
						})

						session := helpers.CustomCF(helpers.CFEnv{WorkingDirectory: runningDir}, PushCommandName)
						Eventually(session).Should(helpers.SayPath(`path:\s+%s`, dir))
						Eventually(session).Should(Exit(0))
					})
				})
			})
		})
	})
})

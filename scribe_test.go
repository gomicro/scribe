package scribe

import (
	"testing"

	"github.com/alecthomas/assert"
	"github.com/gomicro/penname"
)

func TestScribe(t *testing.T) {
	t.Run("Describe", func(t *testing.T) {
		mockWrite := penname.New()
		s := NewScribe(mockWrite, DefaultTheme)

		s.BeginDescribe("Header 1")
		{
			s.BeginDescribe("Header 1.1")
			{
				s.BeginDescribe("Header 1.1.1")
				s.EndDescribe()
			}
			s.EndDescribe()

			s.BeginDescribe("Header 1.2")
			s.EndDescribe()
		}
		s.EndDescribe()

		s.BeginDescribe("Header 2")
		s.EndDescribe()

		a := string(mockWrite.Written())
		e := "\nHeader 1\n\n  Header 1.1\n\n    Header 1.1.1\n\n  Header 1.2\n\nHeader 2\n"
		assert.Equal(t, e, a)
	})

	t.Run("Full", func(t *testing.T) {
		mockWrite := penname.New()
		s := NewScribe(mockWrite, DefaultTheme)

		s.BeginDescribe("Organization")

		{
			s.BeginDescribe("Permissions")
			s.EndDescribe()
			s.Done("Enable create private repos")
			s.Done("Enable create public repos")
			s.Done("Base permissions [admin]")
		}

		{
			s.BeginDescribe("Members")
			s.EndDescribe()
			s.Done("Adding John")
			s.Done("Adding Jane")
			s.Done("Adding Jim")
			s.Done("Adding Joe")
		}

		{
			s.BeginDescribe("Teams")
			{
				s.BeginDescribe("Admins")
				s.EndDescribe()
				s.Done("Adding John")
				s.Done("Adding Jane")

				s.BeginDescribe("Developers")
				s.EndDescribe()
				s.Done("Adding Jim")
				s.Done("Adding Joe")
			}
			s.EndDescribe()
		}

		{
			s.BeginDescribe("Repositories")
			{
				s.BeginDescribe("Repo 1")
				s.EndDescribe()
				s.Done("Create repo 1")
				s.Done("Set branch protection")
				s.Done("Set default branch to 'main'")
			}
			{
				s.BeginDescribe("Repo 2")
				s.EndDescribe()
				s.Done("Create repo 2")
				s.Done("Set branch protection")
				s.Done("Set default branch to 'foo'")
			}
			s.EndDescribe()
		}

		s.EndDescribe()

		a := string(mockWrite.Written())
		e := "\nOrganization\n\n  Permissions\n    Enable create private repos\n    Enable create public repos\n    Base permissions [admin]\n\n  Members\n    Adding John\n    Adding Jane\n    Adding Jim\n    Adding Joe\n\n  Teams\n\n    Admins\n      Adding John\n      Adding Jane\n\n    Developers\n      Adding Jim\n      Adding Joe\n\n  Repositories\n\n    Repo 1\n      Create repo 1\n      Set branch protection\n      Set default branch to 'main'\n\n    Repo 2\n      Create repo 2\n      Set branch protection\n      Set default branch to 'foo'\n"
		assert.Equal(t, e, a)
	})

	t.Run("Themed", func(t *testing.T) {
		mockWrite := penname.New()

		theme := &Theme{
			Describe: func(desc string) string {
				return "\033[1;36m" + desc + "\033[0m"
			},
			Done: NoopDecorator,
		}

		s := NewScribe(mockWrite, theme)

		s.BeginDescribe("Organization")

		{
			s.BeginDescribe("Permissions")
			s.EndDescribe()
			s.Done("Enable create private repos")
			s.Done("Enable create public repos")
			s.Done("Base permissions [admin]")
		}

		{
			s.BeginDescribe("Members")
			s.EndDescribe()
			s.Done("Adding John")
			s.Done("Adding Jane")
			s.Done("Adding Jim")
			s.Done("Adding Joe")
		}

		{
			s.BeginDescribe("Teams")
			{
				s.BeginDescribe("Admins")
				s.EndDescribe()
				s.Done("Adding John")
				s.Done("Adding Jane")

				s.BeginDescribe("Developers")
				s.EndDescribe()
				s.Done("Adding Jim")
				s.Done("Adding Joe")
			}
			s.EndDescribe()
		}

		{
			s.BeginDescribe("Repositories")
			{
				s.BeginDescribe("Repo 1")
				s.EndDescribe()
				s.Done("Create repo 1")
				s.Done("Set branch protection")
				s.Done("Set default branch to 'main'")
			}
			{
				s.BeginDescribe("Repo 2")
				s.EndDescribe()
				s.Done("Create repo 2")
				s.Done("Set branch protection")
				s.Done("Set default branch to 'foo'")
			}
			s.EndDescribe()
		}

		s.EndDescribe()

		a := string(mockWrite.Written())
		e := "\n\x1b[1;36mOrganization\x1b[0m\n\n  \x1b[1;36mPermissions\x1b[0m\n    Enable create private repos\n    Enable create public repos\n    Base permissions [admin]\n\n  \x1b[1;36mMembers\x1b[0m\n    Adding John\n    Adding Jane\n    Adding Jim\n    Adding Joe\n\n  \x1b[1;36mTeams\x1b[0m\n\n    \x1b[1;36mAdmins\x1b[0m\n      Adding John\n      Adding Jane\n\n    \x1b[1;36mDevelopers\x1b[0m\n      Adding Jim\n      Adding Joe\n\n  \x1b[1;36mRepositories\x1b[0m\n\n    \x1b[1;36mRepo 1\x1b[0m\n      Create repo 1\n      Set branch protection\n      Set default branch to 'main'\n\n    \x1b[1;36mRepo 2\x1b[0m\n      Create repo 2\n      Set branch protection\n      Set default branch to 'foo'\n"
		assert.Equal(t, e, a)
	})
}

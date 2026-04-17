package scribe

import (
	"errors"
	"strings"
	"sync"
	"testing"

	"github.com/alecthomas/assert"
	"github.com/gomicro/penname"
)

func TestScribe(t *testing.T) {
	t.Parallel()

	t.Run("Describe", func(t *testing.T) {
		t.Parallel()
		mockWrite := penname.New()
		s, err := NewScribe(mockWrite, DefaultTheme())
		assert.NoError(t, err)

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
		t.Parallel()
		mockWrite := penname.New()
		s, err := NewScribe(mockWrite, DefaultTheme())
		assert.NoError(t, err)

		s.BeginDescribe("Organization")

		{
			s.BeginDescribe("Permissions")
			s.EndDescribe()
			s.Print("Enable create private repos")
			s.Print("Enable create public repos")
			s.Print("Base permissions [admin]")
		}

		{
			s.BeginDescribe("Members")
			s.EndDescribe()
			s.Print("Adding John")
			s.Print("Adding Jane")
			s.Print("Adding Jim")
			s.Print("Adding Joe")
		}

		{
			s.BeginDescribe("Teams")
			{
				s.BeginDescribe("Admins")
				s.EndDescribe()
				s.Print("Adding John")
				s.Print("Adding Jane")

				s.BeginDescribe("Developers")
				s.EndDescribe()
				s.Print("Adding Jim")
				s.Print("Adding Joe")
			}
			s.EndDescribe()
		}

		{
			s.BeginDescribe("Repositories")
			{
				s.BeginDescribe("Repo 1")
				s.EndDescribe()
				s.Print("Create repo 1")
				s.Print("Set branch protection")
				s.Print("Set default branch to 'main'")
			}
			{
				s.BeginDescribe("Repo 2")
				s.EndDescribe()
				s.Print("Create repo 2")
				s.Print("Set branch protection")
				s.Print("Set default branch to 'foo'")
			}
			s.EndDescribe()
		}

		{
			s.BeginDescribe("Output")
			{
				s.PrintLines(strings.NewReader("Output line 1\nOutput line 2\n"))
			}
			s.EndDescribe()
		}

		{
			s.BeginDescribe("Errors")
			{
				s.Error(errors.New("something went wrong"))
			}
			s.EndDescribe()
		}

		s.EndDescribe()

		a := string(mockWrite.Written())
		e := "\nOrganization\n\n  Permissions\n    Enable create private repos\n    Enable create public repos\n    Base permissions [admin]\n\n  Members\n    Adding John\n    Adding Jane\n    Adding Jim\n    Adding Joe\n\n  Teams\n\n    Admins\n      Adding John\n      Adding Jane\n\n    Developers\n      Adding Jim\n      Adding Joe\n\n  Repositories\n\n    Repo 1\n      Create repo 1\n      Set branch protection\n      Set default branch to 'main'\n\n    Repo 2\n      Create repo 2\n      Set branch protection\n      Set default branch to 'foo'\n\n  Output\n      Output line 1\n      Output line 2\n\n\n  Errors\n      something went wrong\n"
		assert.Equal(t, e, a)
	})

	t.Run("Themed", func(t *testing.T) {
		t.Parallel()
		mockWrite := penname.New()

		theme := &Theme{
			Describe: func(desc string) string {
				return "\033[1;36m" + desc + "\033[0m"
			},
			Print: NoopDecorator,
			Error: func(err error) string {
				return "Error: " + err.Error()
			},
		}

		s, err := NewScribe(mockWrite, theme)
		assert.NoError(t, err)

		s.BeginDescribe("Organization")

		{
			s.BeginDescribe("Permissions")
			s.EndDescribe()
			s.Print("Enable create private repos")
			s.Print("Enable create public repos")
			s.Print("Base permissions [admin]")
		}

		{
			s.BeginDescribe("Members")
			s.EndDescribe()
			s.Print("Adding John")
			s.Print("Adding Jane")
			s.Print("Adding Jim")
			s.Print("Adding Joe")
		}

		{
			s.BeginDescribe("Teams")
			{
				s.BeginDescribe("Admins")
				s.EndDescribe()
				s.Print("Adding John")
				s.Print("Adding Jane")

				s.BeginDescribe("Developers")
				s.EndDescribe()
				s.Print("Adding Jim")
				s.Print("Adding Joe")
			}
			s.EndDescribe()
		}

		{
			s.BeginDescribe("Repositories")
			{
				s.BeginDescribe("Repo 1")
				s.EndDescribe()
				s.Print("Create repo 1")
				s.Print("Set branch protection")
				s.Print("Set default branch to 'main'")
			}
			{
				s.BeginDescribe("Repo 2")
				s.EndDescribe()
				s.Print("Create repo 2")
				s.Print("Set branch protection")
				s.Print("Set default branch to 'foo'")
			}
			s.EndDescribe()
		}

		{
			s.BeginDescribe("Errors")
			{
				s.Error(errors.New("something went wrong"))
			}
			s.EndDescribe()
		}

		s.EndDescribe()

		a := string(mockWrite.Written())
		e := "\n\x1b[1;36mOrganization\x1b[0m\n\n  \x1b[1;36mPermissions\x1b[0m\n    Enable create private repos\n    Enable create public repos\n    Base permissions [admin]\n\n  \x1b[1;36mMembers\x1b[0m\n    Adding John\n    Adding Jane\n    Adding Jim\n    Adding Joe\n\n  \x1b[1;36mTeams\x1b[0m\n\n    \x1b[1;36mAdmins\x1b[0m\n      Adding John\n      Adding Jane\n\n    \x1b[1;36mDevelopers\x1b[0m\n      Adding Jim\n      Adding Joe\n\n  \x1b[1;36mRepositories\x1b[0m\n\n    \x1b[1;36mRepo 1\x1b[0m\n      Create repo 1\n      Set branch protection\n      Set default branch to 'main'\n\n    \x1b[1;36mRepo 2\x1b[0m\n      Create repo 2\n      Set branch protection\n      Set default branch to 'foo'\n\n  \x1b[1;36mErrors\x1b[0m\n      Error: something went wrong\n"
		assert.Equal(t, e, a)
	})
}

func TestPrintf(t *testing.T) {
	t.Parallel()

	mockWrite := penname.New()
	s, err := NewScribe(mockWrite, DefaultTheme())
	assert.NoError(t, err)

	s.BeginDescribe("Section")
	s.Printf("item %d of %d", 1, 3)
	s.EndDescribe()

	a := string(mockWrite.Written())
	e := "\nSection\n    item 1 of 3\n"
	assert.Equal(t, e, a)
}

func TestErrorf(t *testing.T) {
	t.Parallel()

	mockWrite := penname.New()
	s, err := NewScribe(mockWrite, DefaultTheme())
	assert.NoError(t, err)

	s.BeginDescribe("Section")
	s.Errorf("failed after %d attempts", 3)
	s.EndDescribe()

	a := string(mockWrite.Written())
	e := "\nSection\n    failed after 3 attempts\n"
	assert.Equal(t, e, a)
}

func TestChild(t *testing.T) {
	t.Parallel()

	t.Run("Sequential", func(t *testing.T) {
		t.Parallel()

		mockWrite := penname.New()
		s, err := NewScribe(mockWrite, DefaultTheme())
		assert.NoError(t, err)

		s.BeginDescribe("Parent")

		child1 := s.Child("Child 1")
		child1.Print("item a")
		child1.EndDescribe()

		child2 := s.Child("Child 2")
		child2.Print("item b")
		child2.EndDescribe()

		s.EndDescribe()

		a := string(mockWrite.Written())
		e := "\nParent\n\n  Child 1\n      item a\n\n  Child 2\n      item b\n"
		assert.Equal(t, e, a)
	})

	t.Run("Concurrent", func(t *testing.T) {
		t.Parallel()

		mockWrite := penname.New()
		s, err := NewScribe(mockWrite, DefaultTheme())
		assert.NoError(t, err)

		s.BeginDescribe("Parent")

		items := []string{"Alpha", "Beta", "Gamma"}
		children := make([]Scriber, len(items))
		for i, name := range items {
			children[i] = s.Child(name)
		}

		var wg sync.WaitGroup
		for i, child := range children {
			wg.Add(1)
			go func(c Scriber, label string) {
				defer wg.Done()
				c.Print("processing " + label)
				c.EndDescribe()
			}(child, items[i])
		}
		wg.Wait()

		s.EndDescribe()

		output := string(mockWrite.Written())
		assert.Contains(t, output, "\n  Alpha\n      processing Alpha\n")
		assert.Contains(t, output, "\n  Beta\n      processing Beta\n")
		assert.Contains(t, output, "\n  Gamma\n      processing Gamma\n")
	})
}
